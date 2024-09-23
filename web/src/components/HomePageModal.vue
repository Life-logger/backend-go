<template>
  <div class="modal-overlay" @click="close">
    <div class="modal-container" @click.stop>
      
      <!-- 모달 헤더 부분 -->
      <div class="log-header">
        <!-- 선택된 로그 색상 -->
        <div class="modal-log-color">
          <div class="log-color"></div>
        </div>
        <!-- 로그 제목 및 시간 -->
        <div class="modal-log-title"></div>
        <!-- 핀 버튼과 수정 버튼 -->
        <div class="modal-two-button">
          <div class="modal-pin-button"></div>
          <div class="modal-edit-button"></div>
        </div>
      </div>

      <!-- 이미지 업로드 버튼 -->
      <div id="image-upload-container" @click.stop="triggerFileInput">
        <input
          type="file"
          id="fileInput"
          accept="image/*"
          @change="handleFileChange"
          ref="fileInput"
          style="display: none;"
        />
        <label for="fileInput" class="custom-upload-button" :style="{ backgroundImage: imageSrc ? `url(${imageSrc})` : '' }"></label>
      </div>

      <!-- 모달 콘텐츠 -->
      <div class="modal-content" @click.stop>
        <!-- 이미지 미리보기 -->
        <div class="image-preview" v-if="imageSrc">
          <img :src="imageSrc" alt="Preview Image" />
        </div>

        <!-- 텍스트 입력창 -->
        <textarea v-model="modals[activeModal].text" placeholder="내용을 입력하세요."></textarea>

        <!-- 이전, 다음 버튼 -->
        <div class="modal-navigation">
          <button @click="prevModal" :disabled="activeModal === 0" class="prev-button"></button>
          <button @click="nextModal" :disabled="activeModal === modals.length - 1" class="next-button"></button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HomePageModal',
  data() {
    return {
      activeModal: 0,
      modals: [
        { text: '' }, // 모달의 텍스트 데이터
        { text: '' }, // 모달의 텍스트 데이터
        // 필요한 만큼 추가
      ],
      imageSrc: null,
      isFileInputTriggered: false,
    };
  },
  methods: {
    close() {
      this.$emit('close');
    },
    prevModal() {
      if (this.activeModal > 0) {
        this.activeModal--;
      }
    },
    nextModal() {
      if (this.activeModal < this.modals.length - 1) {
        this.activeModal++;
      }
    },
    handleFileChange(event) {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          this.imageSrc = e.target.result;
          this.isFileInputTriggered = false;
        };
        reader.readAsDataURL(file);
      } else {
        this.isFileInputTriggered = false;
      }
    },
    triggerFileInput() {
      if (!this.isFileInputTriggered) {
        this.isFileInputTriggered = true;
        if (this.$refs.fileInput) {
          this.$refs.fileInput.click();
        }
      }
    }
  }
};
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-container {
  background-color: #ffffff;
  padding: 20px;
  border-radius: 8px;
  width: 358px;
  height: 504px;
  overflow-y: auto;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
  justify-content: space-between;
}

.log-header {
  display: flex;
  justify-content: space-between;
  width: 100%;
}

.modal-log-title {
  width: 88px;
  height: 38px;
  background-color: #ccc;
}

.modal-log-color {
  width: 20px;
  height: 38px;
  background-color: #000;
  display: flex;
  align-items: center;
  justify-content: center;
}

.log-color {
  width: 15px;
  height: 15px;
  background-color: #BEC6A0;
  border-radius: 50%;
}

.modal-two-button{
  width:24px;
  height:52px;
  background-color: #ccc;
}

.modal-pin-button, .modal-edit-button{
  width:24px;
  height:24px;
  background-color: #fff;
  margin : 1px 0px;
}


#image-upload-container { /* 이미지 버튼이 들어갈 영역 설정 */
  position: relative;
  width: 110px;
  height: 79px;
}

.custom-upload-button { /* 이미지 버튼 커스텀 */
  display: block;
  width: 100%;
  height: 100%;
  background-image: url('@/assets/Image_upload_button_icon.png');
  background-size: cover;
  cursor: pointer;
  position: absolute; 
  top: 0;
  left: 0;
}

input[type="file"] { /* 이미지 업로드 버튼의 '파일선택'을 숨기기*/
  display: none;
}

.image-preview img {
  max-width: 100%;
  max-height: 200px;
}

textarea {
  width: 318px;
  height: 313px;
  padding: 10px;
  margin: 8px;
  border-radius: 4px;
  border: 1px solid #6eb268;
  box-sizing: border-box;
}

.textarea-container {
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  flex-grow: 1;
}

.modal-navigation {
  display: flex;
  justify-content: space-between;
  width: 100%;
}

.modal-navigation button {
  background-color: #ffffff;
  border: none;
  cursor: pointer;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  background-size: cover;
}

.prev-button {
  left: 22.5%;
  background-image: url('@/assets/Modal_button1_icon.png');
}

.next-button {
  right: 22.5%;
  background-image: url('@/assets/Modal_button2_icon.png');
}

.prev-button, .next-button {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  z-index: 1;
  background-color: #ccc;
  border: none;
  cursor: pointer;
  border-radius: 50%;
  width: 28px;
  height: 28px;
}
</style>