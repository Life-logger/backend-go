TimerCategoryList.vue
<template>
  <div class="timer-category-list">
    <div
      v-for="(category, index) in categories"
      :key="category.id"
      :style="{ backgroundColor: category.color }"
      class="category-item"
    >
      <h3>{{ category.name }}</h3>
      <p>{{ formatTime(category.timer) }}</p>

      <!-- 타이머 시작/정지 버튼 -->
      <button @click="toggleTimer(index)">
        <img
          :src="getIconSrc(category.timerRunning ? `Timer_Stop_icon.png` : `Timer_Start_icon.png`)"
          :alt="category.timerRunning ? 'Stop Timer' : 'Start Timer'"
        />
      </button>

      <!-- 타이머 초기화 버튼 -->
      <button @click="resetTimer(index)">
        <img 
          :src="getIconSrc('Reset_Timer_icon.png')" 
          :alt="`Reset Timer`"
          />
      </button>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';

export default {
  name: 'TimerCategoryList',
  computed: {
    ...mapGetters(['getCategories']),
    categories() {
      return this.getCategories;
    }
  },
  methods: {
    ...mapActions(['updateTimer', 'resetTimer']),
    toggleTimer(index) {
      this.$emit('toggle-timer', index); // 상위 컴포넌트에서 관리하도록 이벤트 전송
    },
    resetTimer(index) {
      this.$emit('reset-timer', index); // 상위 컴포넌트에서 관리하도록 이벤트 전송
    },
    formatTime(seconds) {
      const hours = String(Math.floor(seconds / 3600)).padStart(2, '0');
      const minutes = String(Math.floor((seconds % 3600) / 60)).padStart(2, '0');
      const secs = String(seconds % 60).padStart(2, '0');
      return `${hours}:${minutes}:${secs}`;
    },
    getIconSrc(image) { 
      const url = new URL(`@/assets/${image}`, import.meta.url).href;
      console.log(url); // 로그 추가
      return url;
    }
  }
};
</script>

<style scoped>
.timer-category-list {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  overflow-x: auto;
  scrollbar-width: none;
  position: absolute;
  align-content: flex-start;
  column-gap: 0;
  width: 360px;
  height: 577px;
  left: 15px;
  top: 171px;
  background: rgba(218, 225, 244, 0.5);
}  

.category-timer {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border: none;
  margin-left: 15px;
  margin-top: 15px;
  padding: 0;
  gap: 10px;
  width: 100px;
  height: 150px;
  box-shadow: 0px 2px 5px rgba(116, 116, 116, 0.3);
  border-radius: 16px;
}

.category-timer p {
  font-size: 2rem;
  margin: 0px;
  padding: 0px;
  gap: 0;
  justify-content: center;
  width: 77px;
  height: 20px;
  font-family: 'Pretendard';
  font-weight: 700;
  font-size: 17px;
  line-height: 20px;
  text-align: center;
  color: #303030;
}

.category-timer button {
  background: transparent;
  border: none;
  cursor: pointer;
  margin: 0px;
  padding: 0px;
  gap: 0;
  width: 40px;
  height: 40px;
}
</style>