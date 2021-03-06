import {FormCard} from "../../Base/Base";
import {Button} from "react-bootstrap";
import React, {useState} from "react";
import {FORM_PRE_ENROLL_CODE, FORM_PRE_ENROLL_DETAILS, usePreEnrollment} from "./PreEnrollmentFlow";
import InputGroup from "react-bootstrap/InputGroup";
import Form from "react-bootstrap/Form";
import {PHONE_NUMBER_MAX} from "../../Login/EnterPhoneNumberComponent";
import "./EnterPreEnrollment.scss"
import {BaseFormCard} from "../../components/BaseFormCard";

export function PreEnrollmentCode(props) {
    const {goBack} = usePreEnrollment()
    // return (
    //     <FormCard onBack={() => {
    //         goBack()
    //     }} content={<EnterPreEnrollmentContent/>} title={"Verify Recipient"}/>
    // );
    return (

        <BaseFormCard title={"Verify Recipient"}>
            <EnterPreEnrollmentContent/>
        </BaseFormCard>
    )
}

function EnterPreEnrollmentContent(props) {
    const {state, goNext} = usePreEnrollment()
    const [phoneNumber, setPhoneNumber] = useState(state.mobileNumber)
    const [enrollCode, setEnrollCode] = useState(state.enrollCode)

    const handlePhoneNumberOnChange = (e) => {
        if (e.target.value.length <= PHONE_NUMBER_MAX) {
            setPhoneNumber(e.target.value)
        }
    }

    const handleEnrollCodeOnChange = (e) => {
        if (e.target.value.length <= 5) {
            setEnrollCode(e.target.value)
        }
    }
    return (
        <div className="enroll-code-container">
            <h4 className="title text-center">Enter Mobile Number & Pre Enrolment Code</h4>
            <div className={"input-container"}>
                <div className="divOuter">
                    <div className="divInner">

                        <Form.Group>
                            <Form.Control type="text" placeholder="+91-XXXXXXXXX" tabIndex="1" value={phoneNumber} onChange={handlePhoneNumberOnChange}/>
                            <Form.Control type="text" placeholder="XXXXX" tabIndex="1" value={enrollCode} onChange={handleEnrollCodeOnChange}/>
                            {/*<input id="otp" type="text" className="otp" tabIndex="2" maxLength="5"*/}
                            {/*       value={enrollCode}*/}
                            {/*       onChange={handleEnrollCodeOnChange}*/}
                            {/*       placeholder=""/>*/}
                        </Form.Group>

                    </div>
                </div>
            </div>
            <Button variant="outline-primary" className="action-btn" onClick={() => {
                goNext(FORM_PRE_ENROLL_CODE, FORM_PRE_ENROLL_DETAILS, {
                    mobileNumber: phoneNumber,
                    enrollCode: enrollCode
                })
            }}>CONFIRM</Button>
        </div>
    );
}
