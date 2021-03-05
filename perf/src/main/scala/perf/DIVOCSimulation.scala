package perf

import io.gatling.core.Predef._
import io.gatling.core.scenario.Simulation
import io.gatling.http.Predef._

import scala.concurrent.duration._
import scala.util.Random

class DIVOCSimulation extends Simulation {
  private val env: String = System.getenv.getOrDefault("targetEnv", "dev")
  var baseUrl = System.getenv.getOrDefault("baseUrl", "http://localhost:81")
  val token = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJIbldDQnE0MTF6YWpKWjg2RUN6WGptd0ptTDljQU5NaFVFZkdTSEdVcnpzIn0.eyJleHAiOjE2MzIxMzkxNzEsImlhdCI6MTYxNDg1OTE3MiwiYXV0aF90aW1lIjoxNjE0ODU5MTcxLCJqdGkiOiJlOTg3NmVkNC05YjE0LTRiZTQtODYwYy01NzgxOTkyYWFlOWUiLCJpc3MiOiJodHRwOi8vZGl2b2MueGl2LmluL2F1dGgvcmVhbG1zL2Rpdm9jIiwiYXVkIjpbInJlYWxtLW1hbmFnZW1lbnQiLCJhY2NvdW50IiwiZmFjaWxpdHktYWRtaW4tcG9ydGFsIl0sInN1YiI6IjZhYjY0NDQ1LTBmN2YtNGEwYy04ZWMwLTRjMDUzMDg3Y2VmMyIsInR5cCI6IkJlYXJlciIsImF6cCI6ImZhY2lsaXR5LWFwaSIsIm5vbmNlIjoiZWVlODJlYzctNjRiNC00NTk5LTg4MmEtMGRjODc5NWQ1MGQ1Iiwic2Vzc2lvbl9zdGF0ZSI6Ijc4NjUzNjdjLTdjM2UtNDQxMC1iOWJlLWQxNDEzMThiNmY2MyIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOlsiaHR0cHM6Ly9kaXZvYy54aXYuaW4iLCJodHRwOi8vbG9jYWxob3N0OjUwMDAiLCIqIiwiaHR0cHM6Ly9kaXZvYy1hcHAueGl2LmluIiwiaHR0cDovL2xvY2FsaG9zdDozMDAwIl0sInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJvZmZsaW5lX2FjY2VzcyIsInVtYV9hdXRob3JpemF0aW9uIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsicmVhbG0tbWFuYWdlbWVudCI6eyJyb2xlcyI6WyJ2aWV3LXVzZXJzIiwicXVlcnktZ3JvdXBzIiwicXVlcnktdXNlcnMiXX0sImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfSwiZmFjaWxpdHktYWRtaW4tcG9ydGFsIjp7InJvbGVzIjpbImZhY2lsaXR5LXN0YWZmIl19fSwic2NvcGUiOiJvcGVuaWQgcHJvZmlsZSBlbWFpbCIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwicHJlZmVycmVkX3VzZXJuYW1lIjoiMTExMTExMTEyNSIsImZhY2lsaXR5X2NvZGUiOiJNQUExMyJ9.cRdvv-4xYr-fSs5xr1sOuRFN6O091WgmnSNwxjBSxbdM5scGaOwLlcrTsLXi1abKVztwQCOlBPzSz0kcY8baEn1Jo_Q4igojfxaLkNJ3fITbxX2cI_3K0ABSjGybPy_qbbxZuciJmQpFV-nMbAtyiX2aAJw1uJpOijB3VWTM5BIDLLMz934PUn8nOYYG_IGImjc8AiJbRdbLL2ccrTRkS5LfbHKBym2csUioswYwxyyc12SGe1Xi_bcLlSLY28vd3824zSHV23VtyyJs0NRTgNIJtJinukHIElyDcgY7HoRXVLJtCKdJqmmBN8DxsmdJFvs_kkUExerwWDU2ZNtikw"
  println(s"Using env ${baseUrl}")

  val httpProtocol = http
    .warmUp(baseUrl)
    .baseUrl(baseUrl)
    .acceptHeader("application/json")
    .contentTypeHeader("application/json")

  val headers = Map(
    "accept" -> "application/json",
    "Content-Type" -> "application/json",
    "Authorization" -> s"Bearer ${token}"
  )

//  val requestBody = RawFileBody("certificate-request.json")

  val feeder = Iterator.continually(Map("id" -> Random.nextInt()))

  val req = http("Create certificate")
    .post(s"${baseUrl}/divoc/api/v1/certify")
    .headers(headers)
    .body(StringBody("""[
                       |  {
                       |        "preEnrollmentCode": "118479${r}",
                       |        "recipient": {
                       |            "contact": [
                       |                "tel:${r}"
                       |            ],
                       |            "dob": "1995-03-01",
                       |            "gender": "Male",
                       |            "identity": "did:in.gov.uidai.aadhaar:234234234342",
                       |            "name": "Suresh ${r}",
                       |            "nationality": "Suresh P",
                       |            "address": {
                       |                "addressLine1": "asd",
                       |                "district": "asd",
                       |                "state": "asd",
                       |                "pincode": "asd"
                       |            }
                       |        },
                       |        "vaccination": {
                       |            "batch": "12344312",
                       |            "date": "2020-12-02T09:44:03.802Z",
                       |            "effectiveStart": "2020-12-02",
                       |            "effectiveUntil": "2020-12-02",
                       |            "manufacturer": "string",
                       |            "name": "COVID-19",
                       |            "dose": 1,
                       |            "totalDoses": 2
                       |        },
                       |        "vaccinator": {
                       |            "name": "Nayan A"
                       |        },
                       |        "facility": {
                       |            "name": "Awesome Hospital",
                       |            "address": {
                       |                "addressLine1": "asd",
                       |                "district": "asd",
                       |                "state": "asd",
                       |                "pincode": "asd"
                       |            }
                       |        }
                       |    }
                       |]""".stripMargin))
    .check(status.in(200))


  val scn = scenario("Certificate")
    .feed(feeder)
    .forever() {
      pace(1 seconds)
        .exec({s=>
          s.set("r", Random.nextInt().toString)
        }
        )
        .exec(req)

    }


//  setUp(scn.inject(atOnceUsers(1)))

  setUp(scn.inject(
    atOnceUsers(1),
    rampUsers(50).during((300/5).seconds)
  )).protocols(httpProtocol).assertions(
    global.responseTime.max.lt(800),
    global.successfulRequests.percent.gt(98)
  ).maxDuration(8 hours)
}
