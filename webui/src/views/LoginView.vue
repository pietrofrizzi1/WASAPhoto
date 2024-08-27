<!-- src/views/LoginView.vue -->
<template>
    <div class="login-container">
      <h1>Login</h1>
      <form @submit.prevent="handleLogin">
        <div>
          <label for="username">Username:</label>
          <input type="text" v-model="username" id="username" />
        </div>
        <button type="submit">Login</button>
        <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
      </form>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { login } from '../services/authService';
  
  const username = ref('');
  
  const errorMessage = ref('');
  
  const router = useRouter();
  
  const handleLogin = async () => {
    errorMessage.value = '';
    try {
      // Assicurati che login() sia definito correttamente e gestisca le credenziali
      const response = await login(username.value);
      console.log('Login riuscito:', response); // Per il debug
      router.push('/session'); // Reindirizzamento
    } catch (error) {
      console.error('Errore di login:', error); // Log dell'errore
      errorMessage.value = error.message;
    }
  };
  </script>
  
  <style scoped>
  .login-container {
    max-width: 400px;
    margin: auto;
    padding: 1em;
    border: 1px solid #ccc;
    border-radius: 5px;
  }
  
  .error-message {
    color: red;
  }
  </style>
  