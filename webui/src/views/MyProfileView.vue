<script>
import Comments from "../components/Comments.vue";

export default {
    components: { Comments },
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
					let response = await this.$axios.put("/users/" + this.username + "/photos/" + Math.floor(Math.random() * 10000), this.images, {
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
                let response = await this.$axios.get("/users/" + this.username + "/photos", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.photoList = response.data
                for (let i = 0; i < this.photoList.photos.length; i++) {
                    this.photoList.photos[i].file = 'data:image/*;base64,' + this.photoList.photos[i].file
                     this.photoList.photos[i].comment = ""
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
                let response = await this.$axios.delete("/users/" + this.username + "/photos/" + photoid, {
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
                    let response = await this.$axios.put("/users/" + this.username ,{ username: this.newUsername }, {
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
        async sendComment(username, photoid, comment) {
            if (!comment.trim()) {
                this.errormsg = "Empty comment field.";
                return;
            }

            try {
                await this.$axios.put(`/users/${username}/photos/${photoid}/comments/${Math.floor(Math.random() * 10000)}`, 
                { content: comment }, 
                { headers: { Authorization: `Bearer ${this.token}` } }
                );
                await this.refresh();
            } catch (e) {
                this.handleError(e);
            }
        },

        async openLog(username, photoid) {
            try {
                const response = await this.$axios.get(`/users/${username}/photos/${photoid}/comments`, {
                headers: {
                    Authorization: `Bearer ${this.token}`,
                },
                });
                this.photoComments = response.data;
                const modal = new bootstrap.Modal(document.getElementById('logviewer'));
                modal.show();
            } catch (e) {
                this.handleError(e);
            }
        },

        async likePhoto(username, id) {
            try {
                let response = await this.$axios.put("/users/" + username + "/photos/" + id + "/likes/" + Math.floor(Math.random() * 10000), {}, {
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
        async removeLike(username, id) {
            try {
                let response = await this.$axios.get("/users/" + username + "/photos/" + id + "/likes", {
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
                let response = await this.$axios.delete("/users/" + username + "/photos/" + id + "/likes/" + this.like.likeId, {
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
        async doLogout() {
			localStorage.removeItem("token")
			localStorage.removeItem("username")
			this.$router.push({ path: '/' })
		},
    },
    mounted() {
        this.userProfile()
        this.userPhotos()
    }
}
</script>


<template>
    <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
      <div class="position-sticky pt-3 sidebar-sticky">
        <ul class="nav flex-column">
          <li class="nav-item">
            <RouterLink to="/session" class="nav-link">
              <svg class="feather">
                <use href="/feather-sprite-v4.29.0.svg#image" />
              </svg>
              Stream
            </RouterLink>
          </li>
          <li class="nav-item">
            <div class="nav-link" @click="doLogout">
              <svg class="feather">
                <use href="/feather-sprite-v4.29.0.svg#log-out" />
              </svg>
              Logout
            </div>
          </li>
        </ul>
      </div>
    </nav>
  
    <!-- Main content -->
    <div class="col-md-9 ms-sm-auto col-lg-10 px-4">
      <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
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
            <input
              type="text"
              id="newUsername"
              v-model="newUsername"
              class="form-control"
              placeholder="Insert a new username for your profile..."
            />
            <button class="btn btn-success" @click="changeName">Change username</button>
          </div>
  
          <!-- Photo Upload Section - Styled Like Username Change -->
          <div class="photo-upload">
            <input
              type="file"
              accept="image/*"
              class="form-control"
              @change="uploadFile"
              ref="file"
              placeholder="No image selected"
            />
            <button class="btn btn-primary mt-2" @click="submitFile">Upload Photo</button>
          </div>
  
          <!-- Error Message -->
          <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
  
          <!-- Photo Grid -->
          <div class="photo-stream">
            <div v-for="photo in photoList.photos" :key="photo.id">
              <div class="card shadow-sm">
                <img class="card-img-top" :src="photo.file" alt="Photo" />
                <div class="photo-info">
                  <p class="card-text text-muted">{{ photo.date }}</p>
  
                  <div class="input-group mb-3">
                    <input type="text" v-model="photo.comment" class="form-control" placeholder="Leave a comment" />
                    <button class="btn btn-primary" type="button" @click="sendComment(username, photo.id, photo.comment)">Send</button>
                  </div>
  
                  <!-- Actions -->
                  <div class="photo-actions d-flex justify-content-between align-items-center">
                    <div class="btn-group">
                      <button type="button" class="btn btn-dark" @click="openLog(username, photo.id)">
                        View Comments ({{ photo.commentsCount }})
                      </button>
                      <button v-if="photo.likeStatus == false" class="btn btn-primary" @click="likePhoto(username, photo.id)">
                        Like ({{ photo.likesCount }})
                      </button>
                      <button v-if="photo.likeStatus == true" class="btn btn-danger" @click="removeLike(username, photo.id)">
                        Unlike ({{ photo.likesCount }})
                      </button>
                      <button class="btn btn-outline-danger" @click="deletePhoto(photo.id)">Delete</button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
  
            <!-- Log Modal -->
            <!-- Log Modal -->
            <Comments id="logviewer" :log="photoComments" :token="token"></Comments>

          </div>
        </div>
      </div>
    </div>
  </template>
  





<style>
</style>
<style scoped>
.profile-container {
    display: flex;
}
.container {
  margin-top: 0;
  padding: 0;
  height: 100vh; /* Ocupa tutta l'altezza dello schermo */
  
  display: flex;
  flex-direction: column;
}

.photo-stream {
  flex: 1;
  overflow-y: auto; /* Abilita lo scorrimento verticale */
  display: flex;
  flex-direction: column; /* Dispone le immagini in colonna */
  padding: 10px;
  gap: 20px; /* Spazio tra le foto */
}



/* Card and image styling */
.card {
  border: none;
  margin: 0 auto; /* Centra la card orizzontalmente */
  padding: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  overflow: hidden;
  width: auto; /* Lascia che la larghezza si adatti al contenuto */
  max-width: 90%; /* Limita la larghezza al 90% dello schermo */
}

.card-img-top {
  object-fit: contain; /* L'immagine si adatta alla card mantenendo le proporzioni */
  width: 100%; /* Assicura che l'immagine occupi tutta la larghezza della card */
  height: auto; /* L'altezza si adatta alla proporzione dell'immagine */
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
