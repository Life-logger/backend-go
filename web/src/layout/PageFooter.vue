<template>
    <footer>
        <div class="bottom">
            <router-link
                v-for="(icon, index) in icons"
                :key="icon.name"
                :to="icon.path"
                class="footer-icon"
                :class="{ active: isActive(icon.path) }"
                :style="{ order: index }"
            >
                <img 
                    :src="getIconSrc(icon.img)"
                    :alt="icon.name" 
                    class="icon-image"
                >
            </router-link>
        </div>
    </footer>
</template>

<script>
    export default {
        data(){
            return {
                icons: [
                    { name: 'todo', path: '', img: 'Todo_icon.png' },
                    { name: 'timer', path: '/timer', img: 'Timer_icon.png' },
                    { name: 'today', path: '/', img: 'Today_icon.png' },
                    { name: 'calendar', path: '/calendar', img: 'Calendar_icon.png' },
                    { name: 'info', path: '', img: 'Information_icon.png' }
                ]
            }
        },
        methods: {
            getIconSrc(image) {
                return new URL(`../assets/${image}`, import.meta.url).href;            },
            isActive(path){
                return this.$route.path === path;
            }
        }
    }
</script>

<style scoped>
    footer {
        position: absolute;
        width: 390px;
        height: 83px;
        left: 0px;
        top: 761px;
        background: #FFFFFF;
        z-index: 3;
    }
    .bottom {
        /* Tap_Bar_Buttons */

        /* Auto layout */
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: flex-start;
        padding: 0px 16px;
        gap: 15px;

        position: absolute;
        left: 0%;
        right: 0%;
        top: 0%;
        bottom: 20.19%;
    }
    .icon-image {
        /* Icon */
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        padding: 12px 0px;
        gap: 4px;

        margin: 0 auto;
        width: 38px;
        height: 70px;

        /* Inside auto layout */
        flex: none;
        flex-grow: 0;
        filter:brightness(100%);
    }
    .footer-icon.active .icon-image {
        filter: brightness(50%);
    }
</style>