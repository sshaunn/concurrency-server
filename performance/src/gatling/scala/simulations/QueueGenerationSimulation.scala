package simulations

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class QueueGenerationSimulation extends Simulation {

  val httpProtocol = http
    .baseUrl("http://localhost:8080")
    .header("Content-Type", "application/json")

  val scn = scenario("Generate Queue")
    .exec(
      http("Generate Queue Request")
        .post("/v1/queue/generate")
        .body(StringBody(
          """{
            |    "uuid": "some-connection-uuid",
            |    "maxsize": 10,
            |    "users": [
            |        {
            |            "userid": "userid-user-1",
            |            "firstname": "John",
            |            "lastname": "Doe",
            |            "email": "tt@tt.com"
            |        },
            |        {
            |            "userid": "userid-user-2",
            |            "firstname": "John",
            |            "lastname": "Doe",
            |            "email": "tt@tt.com"
            |        },
            |        {
            |            "userid": "userid-user-3",
            |            "firstname": "Alice",
            |            "lastname": "Smith",
            |            "email": "alice.smith@example.com"
            |        },
            |        {
            |            "userid": "userid-user-4",
            |            "firstname": "Bob",
            |            "lastname": "Johnson",
            |            "email": "bob.johnson@example.com"
            |        },
            |        {
            |            "userid": "userid-user-5",
            |            "firstname": "Emma",
            |            "lastname": "Wilson",
            |            "email": "emma.wilson@example.com"
            |        }
            |    ]
            |}""".stripMargin))
        .check(status.is(202))
    )

  setUp(
    scn.inject(
      atOnceUsers(200),  // Start with 200 users immediately
      constantUsersPerSec(200) during (5.minutes)  // Maintain 200 TPS for 5 minutes
    ).protocols(httpProtocol)
  )
}