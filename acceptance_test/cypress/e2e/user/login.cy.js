import {POST, SUCCESS, ANNIE_ADMIN_TOKEN, BRAD_ADMIN_TOKEN, FORBIDDEN, PETER_USER_TOKEN} from "../constants";
import {options} from "../testUtil";
import spok from "cy-spok";

describe("create user", () => {
    it("create user", () => {
        cy.request(
            options({
                url: "/user",
                method: POST,
                body:{
                    userName: "lifelog",
                    userEmail: "lifelog0826@gmail.com",
                }
            })
        ).then((res) => {
            expect(res.status).to.eq(SUCCESS);
            cy.wrap(res.body).should(
                spok({
                    name: "lifelog",
                    email: "lifelog0826@gmail.com",  
                })
            )
        });
    });
});