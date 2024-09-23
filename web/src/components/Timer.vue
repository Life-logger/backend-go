<template>
  <div class="timer-app">
    <div id="timer-title">
      <h2>{{ formattedDate }}<br>Life Logging</h2>
    </div>

    <!-- 카테고리 추가 버튼 -->
    <div class="timer-category-add">
      <img 
        :src="getIconSrc('NewCategory_icon.png')" 
        alt="NewTimerCategory"
        @click="showTimerModal = true"
      >
    </div>

    <!-- TimerModal 컴포넌트 -->
    <TimerModal
      v-if="showTimerModal"
      @close="closeTimerModal"
      @save-category="saveTimerCategory"
    />

    <!-- TimerCategoryList 컴포넌트 -->
    <TimerCategoryList
      :categories="categories"
      @update-timer="updateTimer" 
      @toggle-timer="toggleTimer"
      @reset-timer="resetTimer"
    />
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';
import { format } from 'date-fns'; // 날짜 포맷팅을 위한 date-fns 라이브러리
import TimerModal from './TimerModal.vue';
import TimerCategoryList from './TimerCategoryList.vue';

export default {
  name: 'TimerPage',
  components: {
    TimerModal,
    TimerCategoryList
  },
  data() {
    return {
      formattedDate: '',
      showTimerModal: false,
    };
  },
  computed: {
    ...mapGetters(['getCategories']),
    categories() {
      return this.getCategories;
    }
  },
  created() {
    const today = new Date();
    this.formattedDate = format(today, 'yyyy.MM.dd');
  },
  methods: {
    ...mapActions(['addCategory', 'updateTimer', 'resetTimer', 'recordTimeBlock']),
    saveTimerCategory(category) {
      this.addCategory({
        id: Date.now(),
        name: category.name.trim(),
        color: category.color || '#FFFFFF',
        timer: 0,
        timerRunning: false,
        timerInterval: null
      });
      this.closeTimerModal();
    },
    closeTimerModal() {
      this.showTimerModal = false;
    },
    toggleTimer(index) {
      const category = this.categories[index];
      if (category.timerRunning) {
        clearInterval(category.timerInterval);
        category.timerRunning = false;

        // 타이머가 멈추면, 기록된 시간을 바탕으로 시간 블록을 채웁니다.
        const timeBlockIndex = this.calculateTimeBlockIndex(); // 시간을 계산하는 함수
        this.recordTimeBlock({ categoryIndex: index, timeBlockIndex });
      } else {
        category.timerInterval = setInterval(() => {
          this.updateTimer({ index, elapsedTime: 1 });
        }, 1000);
        category.timerRunning = true;
      }
    },
    resetTimer(index) {
      clearInterval(this.categories[index].timerInterval);
      this.resetTimer(index);
    },
    calculateTimeBlockIndex() {
      // 현재 시간을 기준으로 블록 인덱스를 계산하는 함수 구현
      const now = new Date();
      const hour = now.getHours();
      const minute = now.getMinutes();
      return hour * 6 + Math.floor(minute / 10); // 예를 들어 10분 단위로 블록을 나눈다면
    },
    getIconSrc(image) {
      return new URL(`../assets/${image}`, import.meta.url).href; // 이미지 경로 반환
    }
  }
};
</script>

<style scoped>
#timer-title {
  position: absolute;
  width: 197px;
  height: 85px;
  left: 16px;
  top: 81px;
}

.timer-category-add {
  pointer-events: auto;
  background-color: transparent;
  border: none;
  padding: 0;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  left: calc(50% - 150px/2 + 106px);
  top: 652px;
  z-index: 1; /* 높은 값으로 설정하여 우선순위 증가 */
}

</style>