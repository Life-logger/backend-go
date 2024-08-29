<script>
    export default{
        name: "Login",
        methods:{
            async clickHandler(){
                const REST_API_KEY = '047f9409d2369cfb8f5e8afa41b2d9eb';
                const REDIRECT_URI = 'http://localhost:22250/login/callback';
                const url = `https://kauth.kakao.com/oauth/authorize?response_type=code&client_id=${REST_API_KEY}&redirect_uri=${REDIRECT_URI}`;

                // const response = await fetch(url);
                window.location.href = url;

            }
        },
        mounted(){
        }
    }
    
</script>

<template>
    <h1>Login Page</h1>
    <img src="/public/img/kakaoLoginBtn.png" id="authorizeBtn" @click="clickHandler">
</template>
