import {POST, SUCCESS} from "../constants";
import {options} from "../testUtil";
import spok from "cy-spok";

describe("create user", () => {
    it("create user", () => {
        cy.request(
            options({
                url: "user",
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
                    userName: "lifelog",
                    userEmail: "lifelog0826@gmail.com",
                })
            )
        });
    });
});