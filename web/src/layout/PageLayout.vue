<template>
    <div>
      <router-view @toggle-footer="toggleFooterOverlay" />
      <div class="overlay" v-if="isFooterOverlayVisible"></div> <!-- 페이지 전체를 덮는 오버레이 -->
      <PageFooter />
    </div>
  </template>
  
  <script>
  import PageFooter from './PageFooter.vue';
  
  export default {
    name: 'PageLayout',
    components: {
      PageFooter,
    },
    data() {
      return {
        isFooterOverlayVisible: false, // Footer 오버레이의 표시 여부
      };
    },
    methods: {
      toggleFooterOverlay(isVisible) {
        this.isFooterOverlayVisible = isVisible; // 모달 열리면 오버레이 활성화
      },
    },
  };
  </script>
  
  <style scoped>
  /* 페이지 전체를 덮는 오버레이 스타일 */
  .overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.5); /* 반투명한 검은색 오버레이 */
    z-index: 999; /* 모달보다 아래에 위치 */
    pointer-events: none; /* 클릭을 막기 위해서 pointer-events 설정 */
  }
  </style>
  