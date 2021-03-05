const jsigs = require('jsonld-signatures');
const config = require('./config/config');
const registry = require('./registry');
const {publicKeyPem, privateKeyPem} = require('./config/keys');
const R = require('ramda');
const {RsaSignature2018} = jsigs.suites;
const {AssertionProofPurpose} = jsigs.purposes;
const {RSAKeyPair} = require('crypto-ld');
const {documentLoaders} = require('jsonld');
const {node: documentLoader} = documentLoaders;
const {contexts} = require('security-context');
const credentialsv1 = require('./credentials.json');
const {vaccinationContext} = require("vaccination-context");
const redis = require('./redis');

const UNSUCCESSFUL = "UNSUCCESSFUL";
const SUCCESSFUL = "SUCCESSFUL";
const DUPLICATE_MSG = "duplicate key value violates unique constraint";

const publicKey = {
  '@context': jsigs.SECURITY_CONTEXT_URL,
  id: 'did:india',
  type: 'RsaVerificationKey2018',
  controller: 'https://example.com/i/india',
  publicKeyPem
};

const customLoader = url => {
  console.log("checking " + url);
  const c = {
    "did:india": publicKey,
    "https://example.com/i/india": publicKey,
    "https://w3id.org/security/v1": contexts.get("https://w3id.org/security/v1"),
    'https://www.w3.org/2018/credentials#': credentialsv1,
    "https://www.w3.org/2018/credentials/v1": credentialsv1,
    "https://cowin.gov.in/credentials/vaccination/v1": vaccinationContext,
  };
  let context = c[url];
  if (context === undefined) {
    context = contexts[url];
  }
  if (context !== undefined) {
    return {
      contextUrl: null,
      documentUrl: url,
      document: context
    };
  }
  if (url.startsWith("{")) {
    return JSON.parse(url);
  }
  console.log("Fallback url lookup for document :" + url)
  return documentLoader()(url);
};


async function signJSON(certificate) {

  const publicKey = {
    '@context': jsigs.SECURITY_CONTEXT_URL,
    id: 'did:india',
    type: 'RsaVerificationKey2018',
    controller: 'https://cowin.gov.in/',
    publicKeyPem
  };
  const controller = {
    '@context': jsigs.SECURITY_CONTEXT_URL,
    id: 'https://cowin.gov.in/',
    publicKey: [publicKey],
    // this authorizes this key to be used for making assertions
    assertionMethod: [publicKey.id]
  };

  const key = new RSAKeyPair({...publicKey, privateKeyPem});

  const signed = await jsigs.sign(certificate, {
    documentLoader: customLoader,
    suite: new RsaSignature2018({key}),
    purpose: new AssertionProofPurpose({
      controller: controller
    }),
    compactProof: false
  });

  console.info("Signed cert " + JSON.stringify(signed));
  return signed;
}

function ageOfRecipient(recipient) {
  if (recipient.age) return recipient.age;
  if (recipient.dob && new Date(recipient.dob).getFullYear() > 1900)
    return "" + (new Date().getFullYear() - new Date(recipient.dob).getFullYear())
  return "";
}

function transformW3(cert, certificateId) {
  const certificateFromTemplate = {
    "@context": [
      "https://www.w3.org/2018/credentials/v1",
      "https://cowin.gov.in/credentials/vaccination/v1"
    ],
    type: ['VerifiableCredential', 'ProofOfVaccinationCredential'],
    credentialSubject: {
      type: "Person",
      id: R.pathOr('', ['recipient', 'identity'], cert),
      refId: R.pathOr('', ['preEnrollmentCode'], cert),
      name: R.pathOr('', ['recipient', 'name'], cert),
      gender: R.pathOr('', ['recipient', 'gender'], cert),
      age: ageOfRecipient(cert.recipient), //from dob
      nationality: R.pathOr('', ['recipient', 'nationality'], cert),
      address: {
        "streetAddress": R.pathOr('', ['recipient', 'address', 'addressLine1'], cert),
        "streetAddress2": R.pathOr('', ['recipient', 'address', 'addressLine2'], cert),
        "district": R.pathOr('', ['recipient', 'address', 'district'], cert),
        "city": R.pathOr('', ['recipient', 'address', 'city'], cert),
        "addressRegion": R.pathOr('', ['recipient', 'address', 'state'], cert),
        "addressCountry": R.pathOr('IN', ['recipient', 'address', 'country'], cert),
        "postalCode": R.pathOr("", ['recipient', 'address', 'pincode'], cert),
      }
    },
    issuer: "https://cowin.gov.in/",
    issuanceDate: new Date().toISOString(),
    evidence: [{
      "id": "https://cowin.gov.in/vaccine/" + certificateId,
      "feedbackUrl": "https://cowin.gov.in/?" + certificateId,
      "infoUrl": "https://cowin.gov.in/?" + certificateId,
      "certificateId": certificateId,
      "type": ["Vaccination"],
      "batch": R.pathOr('', ['vaccination', 'batch'], cert),
      "vaccine": R.pathOr('', ['vaccination', 'name'], cert),
      "manufacturer": R.pathOr('', ['vaccination', 'manufacturer'], cert),
      "date": R.pathOr('', ['vaccination', 'date'], cert),
      "effectiveStart": R.pathOr('', ['vaccination', 'effectiveStart'], cert),
      "effectiveUntil": R.pathOr('', ['vaccination', 'effectiveUntil'], cert),
      "dose": R.pathOr('', ['vaccination', 'dose'], cert),
      "totalDoses": R.pathOr('', ['vaccination', 'totalDoses'], cert),
      "verifier": {
        // "id": "https://nha.gov.in/evidence/vaccinator/" + cert.vaccinator.id,
        "name": R.pathOr('', ['vaccinator', 'name'], cert),
        // "sign-image": "..."
      },
      "facility": {
        // "id": "https://nha.gov.in/evidence/facilities/" + cert.facility.id,
        "name": R.pathOr('', ['facility', 'name'], cert),
        "address": {
          "streetAddress": R.pathOr('', ['facility', 'address', 'addressLine1'], cert),
          "streetAddress2": R.pathOr('', ['facility', 'address', 'addressLine2'], cert),
          "district": R.pathOr('', ['facility', 'address', 'district'], cert),
          "city": R.pathOr('', ['facility', 'address', 'city'], cert),
          "addressRegion": R.pathOr('', ['facility', 'address', 'state'], cert),
          "addressCountry": R.pathOr('IN', ['facility', 'address', 'country'], cert),
          "postalCode": R.pathOr("", ['facility', 'address', 'pincode'], cert)
        },
        // "seal-image": "..."
      }
    }],
    "nonTransferable": "true"
  };
  return certificateFromTemplate;
}

async function signAndSave(certificate, retryCount = 0) {
  const certificateId = "" + Math.floor(1e10 + (Math.random() * 9e10));
  const name = certificate.recipient.name;
  const contact = certificate.recipient.contact;
  const mobile = getContactNumber(contact);
  const preEnrollmentCode = certificate.preEnrollmentCode;
  const currentDose = certificate.vaccination.dose;
  const w3cCertificate = transformW3(certificate, certificateId);
  const signedCertificate = await signJSON(w3cCertificate);
  const signedCertificateForDB = {
    name: name,
    contact: contact,
    mobile: mobile,
    preEnrollmentCode: preEnrollmentCode,
    certificateId: certificateId,
    certificate: JSON.stringify(signedCertificate)
  };
  let resp = await registry.saveCertificate(signedCertificateForDB);
  if (R.pathOr("", ["data", "params", "status"], resp) === UNSUCCESSFUL && R.pathOr("", ["data", "params", "errmsg"], resp).includes(DUPLICATE_MSG)) {
    if (retryCount <= config.CERTIFICATE_RETRY_COUNT) {
      console.error("Duplicate certificate id found, retrying attempt " + retryCount + " of " + config.CERTIFICATE_RETRY_COUNT);
      return await signAndSave(certificate, retryCount + 1)
    } else {
      console.error("Max retry attempted");
      throw new Error(resp.data.params.errmsg)
    }
  }
  if (R.pathOr("", ["data", "params", "status"], resp) === SUCCESSFUL){
    redis.storeKeyWithExpiry(`${preEnrollmentCode}-${currentDose}`, certificateId)
    redis.storeKeyWithExpiry(`${preEnrollmentCode}-cert`, signedCertificateForDB.certificate)
  }
  resp.signedCertificate = {
    ...signedCertificateForDB,
    certificate: signedCertificate,
    meta: certificate["meta"]
  };
  return resp;
}

function getContactNumber(contact) {
  return contact.find(value => /^tel/.test(value)).split(":")[1];
}

module.exports = {
  signAndSave,
  signJSON,
  transformW3,
  customLoader
};
