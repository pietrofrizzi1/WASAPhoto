<script>
export default {
    components: {},
    data: function () {
        return {
            errormsg: null,
            username: "",
            profile: {
                id: 0,
                username: "",
            },
        }
    },
    methods: {
        async doLogin() {
            if (this.username == "") {
                this.errormsg = "Username cannot be empty.";
            } else {
                try {
                    let response = await this.$axios.post("/session", { username: this.username })
                    this.profile = response.data
                    localStorage.setItem("token", this.profile.id);
                    localStorage.setItem("username", this.profile.username);
                    this.$router.push({ path: '/session' })
                } catch (e) {
                    if (e.response && e.response.status === 400) {
                        this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                        this.detailedmsg = null;
                    } else if (e.response && e.response.status === 500) {
                        this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                        this.detailedmsg = e.toString();
                    } else {
                        this.errormsg = e.toString();
                        this.detailedmsg = null;
                    }
                }
            }
        }
    },
    mounted() {
    }
}
</script>

<template>
    <div class="login-container">
        <div class="header">
            <h1>Welcome to WASAPhoto</h1>
        </div>
        <div class="login-form">
            <input type="text" id="username" v-model="username" class="form-control" placeholder="Insert a username to log in WASAPhoto.">
            <button class="btn btn-success login-button" type="button" @click="doLogin">Login</button>
        </div>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    </div>
</template>

<style scoped>
.login-container {
    max-width: 400px;
    margin: 100px auto;
    padding: 30px;
    background-color: #ffffff;
    border-radius: 10px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    text-align: center;
}

.header {
    margin-bottom: 20px;
}

.header h1 {
    font-size: 24px;
    font-weight: bold;
    color: #333;
}

.login-form {
    display: flex;
    flex-direction: column;
}

.form-control {
    margin-bottom: 20px;
    padding: 10px;
    font-size: 16px;
    border: 1px solid #ddd;
    border-radius: 5px;
    transition: border-color 0.2s;
}

.form-control:focus {
    border-color: #007bff;
    outline: none;
    box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
}

.login-button {
    padding: 10px 20px;
    font-size: 16px;
    font-weight: bold;
    color: #ffffff;
    background-color: #28a745;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: background-color 0.3s;
}

.login-button:hover {
    background-color: #218838;
}

.login-button:focus {
    outline: none;
    box-shadow: 0 0 0 2px rgba(40, 167, 69, 0.25);
}
</style>
