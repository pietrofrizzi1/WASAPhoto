<!-- src/views/LoginView.vue -->
<template>
  <div class="login-container">
    <h1>Login</h1>
    <form @submit.prevent="handleLogin">
      <div>
        <label for="username">Username:</label>
        <input type="text" v-model="username" id="username" required />
      </div>
      <button type="submit">Login</button>
      <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
      <p v-if="successMessage" class="success-message">{{ successMessage }}</p>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { doLogin } from '../services/authService'; // Assicurati che il percorso sia corretto

const username = ref('');
const errorMessage = ref('');
const successMessage = ref('');
const router = useRouter();

const handleLogin = async () => {
  errorMessage.value = '';
  successMessage.value = '';
  try {
    
    const response = await doLogin(username.value); // Chiama la funzione doLogin
    
    console.log('Login/Registration successful:', response);
    successMessage.value = 'Login successful!';
    router.push('/session'); // Redirect to another page if needed
  } catch (error) {
    console.error('Login error:', error);
    errorMessage.value = error.message || 'Login failed';
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

.success-message {
  color: green;
}
</style>

