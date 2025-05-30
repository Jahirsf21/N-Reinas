<script setup>
import { useRouter } from 'vue-router'
import { CreateBoard } from '../../wailsjs/go/main/App'
import { ref } from 'vue'

const router = useRouter()
const errorMessage = ref('')

const startGame = () => {
  router.push({ name: 'Game' })
}

</script>

<template>
  <div class="principal-page">
    <h1>Bienvenido</h1>
    <div class="error-message">
      {{ errorMessage }}
    </div>
    <div class="buttons-container">
      <button @click="startGame" class="start-game">Empezar</button>
    </div>
  </div>
</template>

<style scoped>

.principal-page {
  text-align: center;
  padding: 2rem;
}

.buttons-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
  margin-top: 20px;
  margin-bottom: 20px;
}

.start-game {
  padding: 1.5rem;
  background-color: #42b883;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 1.2rem;
  cursor: pointer;
  transition: all 0.3s ease;
  width: 200px;
  text-align: center;
}

.start-game:hover {
  background-color: #3aa876;
  transform: translateY(-3px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.error-message {
  margin-top: 20px;
  color: #ff4757;
  font-weight: bold;
  animation: fadeInOut 3s ease-in-out;
}

@keyframes fadeInOut {
  0% { opacity: 0; }
  20% { opacity: 1; }
  80% { opacity: 1; }
  100% { opacity: 0; }
}
</style>