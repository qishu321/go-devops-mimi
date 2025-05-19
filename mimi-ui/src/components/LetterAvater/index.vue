<template>
    <div class="avatar" :style="{ backgroundColor: bgColor }">
      {{ letter }}
    </div>
  </template>
  
  <script>
  export default {
    name: 'LetterAvatar',
    props: {
      text: {
        type: String,
        default: ''
      }
    },
    computed: {
      letter() {
        return this.text.charAt(0).toUpperCase() || '?';
      },
      bgColor() {
        const colors = ['#409EFF', '#67C23A', '#E6A23C', '#F56C6C', '#909399'];
        const code = this.text.charCodeAt(0) || 0;
        return colors[code % colors.length];
      }
    }
  };
  </script>
  
  <style scoped>
  .avatar {
    width: 32px;
    height: 32px;
    border-radius: 50%;
    color: white;
    font-weight: bold;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 14px;
  }
  </style>