<script>
import LogModal from "../components/Logmodal.vue";

export default {
    components: { LogModal },
    data: function () {
        return {
            errormsg: null,
            username: localStorage.getItem('username'),
            token: localStorage.getItem('token'),
            newUsername: "",
            profile: {
                requestId: 0,
                id: 0,
                username: "",
                followersCount: 0,
                followingCount: 0,
                photoCount: 0,
            },
            photoList: {
                requestUser: 0,
                identifier: 0,
                photos: [
                    {
                        id: 0,
                        userId: 0,
                        file: "",
                        date: "",
                        likesCount: 0,
                        commentsCount: 0,
                        likeStatus: null,
                        comment: "",
                    }
                ],
            },
            user: {
                id: 0,
                username: "",
            },
            comment: "",
            photoComments: {
                requestIdentifier: 0,
                photoIdentifier: 0,
                identifier: 0,
                comments: [
                    {
                        id: 0,
                        userId: 0,
                        photoId: 0,
                        photoOwner: 0,
                        ownerUsername: "",
                        username: "",
                        content: "",
                    }
                ],
            },
        }
    },
    methods: {

        async refresh() {
            await this.userProfile();
            await this.userPhotos();
        },
		async uploadFile() {
			this.images = this.$refs.file.files[0]
		},
		async submitFile() {
			if (this.images === null) {
				this.errormsg = "Please select a file to upload."
			} else {
				try {
					let response = await this.$axios.put("/users/" + this.username + "/photo/" + Math.floor(Math.random() * 10000), this.images, {
						headers: {
							Authorization: "Bearer " + localStorage.getItem("token")
						}
					})
					this.profile = response.data
                    this.refresh()
					this.successmsg = "Photo uploaded successfully."
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
		},
        async userProfile() {
            try {
                let response = await this.$axios.get("/users/" + this.username + "/profile", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.profile = response.data
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
        },
        async userPhotos() {
            try {
                let response = await this.$axios.get("/users/" + this.username + "/photo", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.photoList = response.data
                for (let i = 0; i < this.photoList.photos.length; i++) {
                    this.photoList.photos[i].file = 'data:image/*;base64,' + this.photoList.photos[i].file
                }
            } catch (e) {
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
                    this.detailedmsg = null;
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
                    this.detailedmsg = e.toString();
                } else {
                    this.errormsg = "You haven't posted any photos yet. Go to the home and upload one!";
                    this.detailedmsg = null;
                }
            }
        },
        async deletePhoto(photoid) {
            try {
                let response = await this.$axios.delete("/users/" + this.username + "/photo/" + photoid, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.refresh();
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
        },
        async changeName() {
            if (this.newUsername == "") {
                this.errormsg = "Emtpy username field."
            } else {
                try {
                    let response = await this.$axios.put("/user/" + this.username + "/setusername", { username: this.newUsername }, {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("token")
                        }
                    })
                    this.user = response.data
                    localStorage.setItem("username", this.user.username);
                    this.profile.username = this.user.username;
                    this.username = this.user.username;
                    this.$router.push({ path: '/users/' + this.user.username + '/profile' })
                    this.refresh()
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

        },
        async sendComment(username, photoid,comment) {
            if (comment === "") {
                this.errormsg = "Emtpy comment field."
            } else {
                try {
                    let response = await this.$axios.put("/users/" + username + "/photo/" + photoid + "/comment/" + Math.floor(Math.random() * 10000), { content: comment }, {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("token")
                        }
                    })
                    this.clear = response.data
                    this.refresh()
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
        },
        async openLog(username, photoid) {
            try {
                let response = await this.$axios.get("/users/" + username + "/photo/" + photoid + "/comment", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.photoComments = response.data;
                const modal = new bootstrap.Modal(document.getElementById('logviewer'));
                modal.show();
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
        },
        async likePhoto(username, id) {
            try {
                let response = await this.$axios.put("/users/" + username + "/photo/" + id + "/like/" + Math.floor(Math.random() * 10000), {}, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.clear = response.data
                this.refresh()
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
        },
        async deleteLike(username, id) {
            try {
                let response = await this.$axios.get("/users/" + username + "/photo/" + id + "/like", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.like = response.data
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

            try {
                let response = await this.$axios.delete("/users/" + username + "/photo/" + id + "/like/" + this.like.likeId, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.clear = response.data
                this.refresh()
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
        },
    },
    mounted() {
        this.userProfile()
        this.userPhotos()
    }
}
</script>


<template>
    <div class="profile-container">
        <!-- Sidebar -->
        <nav id="sidebarMenu" class="sidebar bg-light">
            <div class="position-sticky pt-3 sidebar-sticky">
                <h6 class="sidebar-heading text-muted text-uppercase">General</h6>
                <ul class="nav flex-column">
                    <li class="nav-item">
                        <RouterLink to="/session" class="nav-link">
                            <svg class="feather">
                                <use href="/feather-sprite-v4.29.0.svg#home" />
                            </svg>
                            Home
                        </RouterLink>
                    </li>
                </ul>
            </div>
        </nav>

        <!-- Main content -->
        <div class="main-content">
            <header class="profile-header">
                <h1 class="profile-username">{{ profile.username }}</h1>
                <div class="profile-stats">
                    <div class="stat-item">
                        <p class="stat-value">{{ profile.followersCount }}</p>
                        <p class="stat-label">Followers</p>
                    </div>
                    <div class="stat-item">
                        <p class="stat-value">{{ profile.followingCount }}</p>
                        <p class="stat-label">Following</p>
                    </div>
                </div>
            </header>

            <!-- Change Username -->
            <div class="username-change">
                <input type="text" id="newUsername" v-model="newUsername" class="form-control" placeholder="Insert a new username for your profile...">
                <button class="btn btn-success" @click="changeName">Change username</button>
            </div>

            <!-- Photo Upload Section - Styled Like Username Change -->
            <div class="photo-upload">
                <input type="file" accept="image/*" class="form-control" @change="uploadFile" ref="file" placeholder="No image selected">
                <button class="btn btn-primary mt-2" @click="submitFile">Upload Photo</button>
            </div>

            <!-- Error Message -->
            <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

            <!-- Photo Grid -->
            <div class="photo-grid">
                <div class="photo-card" v-for="photo in photoList.photos" :key="photo.id">
                    <img class="photo-image" :src="photo.file" alt="Photo">
                    <div class="photo-info">
                        <RouterLink :to="'/users/' + profile.username + '/profile'" class="username-link">
                            <button class="btn btn-outline-primary">{{ profile.username }}</button>
                        </RouterLink>
                        <p class="photo-date">Photo uploaded on {{ photo.date }}</p>
                        <p class="photo-likes">Likes: {{ photo.likesCount }}</p>
                        <p class="photo-comments">Comments: {{ photo.commentsCount }}</p>
                        
                        <div class="comment-section">
                            <input type="text" id="comment" v-model="photo.comment" class="form-control" placeholder="Comment!">
                            <button class="btn btn-primary" @click="sendComment(username, photo.id, photo.comment)">Send</button>
                        </div>

                        <div class="photo-actions">
                            <button class="btn btn-dark" @click="openLog(username, photo.id)">Comments</button>
                            <button v-if="photo.likeStatus === false" class="btn btn-primary" @click="likePhoto(username, photo.id)">Like</button>
                            <button v-if="photo.likeStatus === true" class="btn btn-danger" @click="deleteLike(username, photo.id)">Unlike</button>
                            <button class="btn btn-outline-danger" @click="deletePhoto(photo.id)">Delete</button>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Log Modal -->
            <LogModal id="logviewer" :log="photoComments" :token="token"></LogModal>
        </div>
    </div>
</template>





<style>
</style>
<style scoped>
.profile-container {
    display: flex;
}

.sidebar {
    flex: 1;
    background-color: #f8f9fa;
    padding: 20px;
    border-right: 1px solid #dee2e6;
}

.sidebar-heading {
    font-weight: bold;
    font-size: 14px;
    margin-bottom: 1rem;
}

.nav-link {
    color: #007bff;
    transition: color 0.2s;
}

.nav-link:hover {
    color: #0056b3;
}

.main-content {
    flex: 3;
    padding: 20px;
}

.profile-header {
    margin-bottom: 2rem;
}

.profile-username {
    font-size: 50px; /* Aumenta la dimensione del testo */
    font-weight: bold; /* Rende il testo in grassetto */
    margin: 0; /* Rimuove il margine inferiore per allineare meglio con le statistiche */
    text-align: center; /* Centra l'username */
}

.profile-stats {
    display: flex;
    justify-content: center; /* Centra le statistiche */
    margin-top: 20px; /* Distanza tra l'username e le statistiche */
}

.stat-item {
    text-align: center;
    margin: 0 30px; /* Spazio tra gli elementi delle statistiche */
}

.stat-value {
    font-size: 40px;
    font-weight: bold;
}

.stat-label {
    font-size: 14px;
    color: #6c757d;
}

.username-change {
    margin-bottom: 2rem;
    display: flex;
    align-items: center;
}

.form-control {
    border-radius: 5px;
    border: 1px solid #ced4da;
    margin-right: 10px;
}

.btn-success {
    border-radius: 5px;
}

.error-message {
    color: #dc3545;
    margin-bottom: 1rem;
}

.photo-upload {
    margin-bottom: 2rem;
    display: flex;
    align-items: center;
}

.photo-grid {
    display: flex;
    flex-wrap: wrap;
    gap: 1rem;
}

.photo-card {
    width: calc(33.333% - 1rem);
    border: 1px solid #dee2e6;
    border-radius: 5px;
    overflow: hidden;
    background-color: #ffffff;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.photo-image {
    width: 100%;
    height: auto;
}

.photo-info {
    padding: 1rem;
}

.username-link {
    margin-bottom: 1rem;
}

.photo-date, .photo-likes, .photo-comments {
    font-size: 14px;
    margin-bottom: 0.5rem;
}

.comment-section {
    margin-bottom: 1rem;
    display: flex;
    align-items: center;
}

.photo-actions {
    display: flex;
    gap: 0.5rem;
    align-items: center;
}

.btn-dark, .btn-primary, .btn-danger, .btn-outline-danger {
    border-radius: 5px;
}

.btn-outline-danger {
    border-color: #dc3545;
    color: #dc3545;
}
</style>
