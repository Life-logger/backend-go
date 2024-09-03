<template>
    <h1>Welcome to the Main Page!</h1>
    <p v-if="isLoggedIn">You are logged in.</p>
    <p v-else>You are not logged in.</p>
    <!-- <a href="/login">Logout</a> -->
</template>

<script>
import {useCookies} from "vue3-cookies";
const {cookies} = useCookies();

const token = cookies.get("accessToken");
export default{
    name: 'Home',
    data(){
        return{
            isLoggedIn:!!token
        }
    },
    mounted(){
        alert(token);
    }
}
</script>