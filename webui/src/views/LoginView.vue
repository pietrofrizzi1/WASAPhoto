<script>
export default {
  data() {
    return {
      errormsg: null,
      username: "",
      profile: {
        id: 0,
        username: "",
      },
    };
  },
  methods: {
    async doLogin() {
      if (this.isUsernameEmpty()) {
        this.setErrorMsg("Username cannot be empty.");
        return;
      }

      try {
        const response = await this.$axios.post("/session", { username: this.username });
        this.handleLoginSuccess(response.data);
      } catch (error) {
        this.handleLoginError(error);
      }
    },

    isUsernameEmpty() {
      return this.username.trim() === "";
    },

    setErrorMsg(message) {
      this.errormsg = message;
      this.detailedmsg = null; // Clear detailed message on new error
    },

    handleLoginSuccess(profileData) {
      this.profile = profileData;
      localStorage.setItem("token", this.profile.id);
      localStorage.setItem("username", this.profile.username);
      this.$router.push({ path: '/session' });
    },

    handleLoginError(error) {
      if (error.response) {
        switch (error.response.status) {
          case 400:
            this.setErrorMsg("Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.");
            break;
          case 500:
            this.setErrorMsg("An internal error occurred. We will be notified. Please try again later.");
            this.detailedmsg = error.toString();
            break;
          default:
            this.setErrorMsg("An unexpected error occurred.");
            this.detailedmsg = error.toString();
        }
      } else {
        this.setErrorMsg("An unexpected error occurred.");
        this.detailedmsg = error.toString();
      }
    },
  },
};
</script>


<template>
    <div class="login-wrapper">
      <div class="login-card">
        <div class="login-card-header">
          <h1 class="login-title">Welcome to WASAPhoto</h1>
          
        </div>
        <div class="login-card-body">
          <input
            type="text"
            id="username"
            v-model="username"
            class="form-input"
            placeholder="Username"
          />
          <button
            class="btn btn-primary login-button"
            type="button"
            @click="doLogin"
          >
            Login
          </button>
        </div>
        <div v-if="errormsg" class="error-message">
          <ErrorMsg :msg="errormsg" />
        </div>
      </div>
    </div>
  </template>
  

  <style scoped>
  /* Wrapper styling */
  .login-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background: linear-gradient(135deg, #ece9e6, #ffffff);
  }
  
  /* Card styling */
  .login-card {
    max-width: 400px;
    width: 100%;
    padding: 30px;
    background-color: #ffffff;
    border-radius: 15px;
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
    text-align: center;
    /* Optional: Add a subtle animation */
    transform: translateY(10px);
    opacity: 0;
    animation: fadeIn 0.5s forwards;
  }
  
  /* Card header styling */
  .login-card-header {
    margin-bottom: 20px;
  }
  
  .login-title {
    font-size: 28px;
    font-weight: 700;
    color: #333;
    margin: 0;
  }
  
  .login-subtitle {
    font-size: 16px;
    color: #666;
  }
  
  /* Input styling */
  .form-input {
    width: 100%;
    padding: 12px 15px;
    margin-bottom: 20px;
    font-size: 16px;
    border: 1px solid #ddd;
    border-radius: 10px;
    transition: border-color 0.3s ease;
  }
  
  .form-input:focus {
    border-color: #007bff;
    outline: none;
    box-shadow: 0 0 0 3px rgba(0, 123, 255, 0.25);
  }
  
  /* Button styling */
  .login-button {
    width: 100%;
    padding: 12px;
    font-size: 18px;
    font-weight: 600;
    color: #ffffff;
    background-color: #007bff;
    border: none;
    border-radius: 10px;
    cursor: pointer;
    transition: background-color 0.3s ease, transform 0.2s ease;
  }
  
  .login-button:hover {
    background-color: #0056b3;
    transform: translateY(-2px);
  }
  
  .login-button:active {
    background-color: #004085;
    transform: translateY(0);
  }
  
  /* Error message styling */
  .error-message {
    margin-top: 20px;
    padding: 10px;
    background-color: #f8d7da;
    color: #721c24;
    border: 1px solid #f5c6cb;
    border-radius: 10px;
    font-size: 14px;
    text-align: left;
  }
  
  /* Keyframes for fade-in animation */
  @keyframes fadeIn {
    from {
      opacity: 0;
      transform: translateY(20px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }
  </style>
  