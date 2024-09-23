<template>
  <div class="modal-overlay" @click="stopAndCloseTimer">
    <div class="modal-content" @click.stop :style="{ backgroundColor: backgroundColor }">
      <h2>{{ category.name }}</h2>
      <p>시작 시간: {{ startTime.toLocaleTimeString() }}</p>
      <p>경과 시간: {{ elapsedTime }}</p>
      <button @click="stopAndCloseTimer">
        <img :src="require(`@/assets/Timer_Stop_icon.png`)" alt="Stop Timer">
      </button>
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  name: 'TimerStartModal',
  props: {
    category: Object,
    backgroundColor: String
  },
  data() {
    return {
      startTime: new Date(),
      timerInterval: null,
      elapsedTime: '00:00:00',
      totalSeconds: 0
    };
  },
  methods: {
    ...mapActions(['updateTimer']),
    stopAndCloseTimer() {
      clearInterval(this.timerInterval);
      this.$emit('close', this.totalSeconds);
    },
    updateElapsedTime() {
      const now = new Date();
      const diff = Math.floor((now - this.startTime) / 1000);
      this.totalSeconds = diff;

      const hours = String(Math.floor(diff / 3600)).padStart(2, '0');
      const minutes = String(Math.floor((diff % 3600) / 60)).padStart(2, '0');
      const seconds = String(diff % 60).padStart(2, '0');

      this.elapsedTime = `${hours}:${minutes}:${seconds}`;
    }
  },
  mounted() {
    this.timerInterval = setInterval(this.updateElapsedTime, 1000);
  },
  beforeUnmount() {
    clearInterval(this.timerInterval);
  }
};
</script>

  
<style scoped>
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: flex-start; /* 상단에 위치시키기 위해 flex-start로 변경 */
  }
  .modal-content {
    display: flex;
    flex-direction: column;
    background: white;
    width: 390px;
    height: 844px;
    margin: 0;
    padding: 0;
    border: none;
    border-radius: 8px;
    text-align: center;
    position: absolute; /* 절대 위치를 사용하여 상단에 배치 */
    top: 0; /* 화면 상단에 고정 */
    align-items: center;
    justify-content: center;
  }
  button {
    background: transparent;
    margin-top: 10px;
    padding: 10px 20px;
    border: none;
    border-radius: 4px;
    color: white;
    cursor: pointer;
  }
</style>
