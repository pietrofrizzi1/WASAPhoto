<template>
    <div class="container">
      <!-- Search Input -->
      <div class="input-group mb-3">
        <input type="text" v-model="searchQuery" class="form-control" placeholder="Search users">
        <button class="btn btn-primary" type="button" @click="searchUsers">Search</button>
      </div>
  
      <!-- Error and Success Messages -->
      <ErrorMsg v-if="errormsg" :msg="errormsg" class="mb-3"></ErrorMsg>
  
      <!-- User List -->
      <div class="list-group">
        <a v-for="user in users" :key="user.id" class="list-group-item list-group-item-action" :href="'/users/' + user.username + '/profile'">
          {{ user.username }}
        </a>
      </div>
    </div>
  </template>
  
  <script>
  import ErrorMsg from "../components/ErrorMsg.vue"; // Assicurati di avere un componente ErrorMsg
  
  export default {
    components: { ErrorMsg },
    data() {
      return {
        searchQuery: "",
        users: [],
        errormsg: null
      };
    },
    methods: {
      async searchUsers() {
        if (this.searchQuery.trim() === "") {
          this.errormsg = "Search field cannot be empty.";
          return;
        }
        try {
          const response = await this.$axios.get(`/users/search/${this.searchQuery}`, {
            headers: {
              Authorization: "Bearer " + localStorage.getItem("token")
            }
          });
          this.users = response.data;
          this.errormsg = null;
        } catch (e) {
          if (e.response && e.response.status === 400) {
            this.errormsg = "Form error, please check all fields and try again.";
          } else if (e.response && e.response.status === 500) {
            this.errormsg = "An internal error occurred. Please try again later.";
          } else {
            this.errormsg = e.toString();
          }
        }
      }
    }
  };
  </script>
  
  <style scoped>
  /* Styling adjustments */
  .container {
    margin-top: 20px;
  }
  .input-group input {
    border-radius: 5px 0 0 5px;
  }
  .input-group button {
    border-radius: 0 5px 5px 0;
  }
  .list-group-item {
    cursor: pointer;
  }
  </style>
  