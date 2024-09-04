<script>
import Comments from "../components/Comments.vue";

export default {
    components: { Comments },
    data: function () {
        return {
            errormsg: null,
            successmsg: null,
            username: localStorage.getItem('username'),
            token: localStorage.getItem('token'),
            profile: {
                requestId: 0,
                id: 0,
                username: "",
                followersCount: 0,
                followingCount: 0,
                photoCount: 0,
                followStatus: null,
                banStatus: null,
                checkIfBanned: null,
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
                    }
                ],
            },
            comment: "",
            follow: {
                followId: 0,
                followedId: 0,
                userId: 0,
                banStatus: 0,
            },
            ban: {
                banId: 0,
                bannedId: 0,
                userId: 0,
            },
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
                        comment: "",
                        username: "",
                        content: "",
                    }
                ],
            },
        }
    },
    methods: {
        async refresh() {
            await this.userProfile()
            await this.userPhotos()
        },
        async userProfile() {
            try {
                let response = await this.$axios.get("/users/" + this.$route.params.username + "/profile", {
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
                let response = await this.$axios.get("/users/" + this.$route.params.username + "/photos", {
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
                    this.errormsg = "User hasen't posted any photos yet."
                    this.detailedmsg = null;
                }
            }
        },
        async banUser(username) {
            try {
                let response = await this.$axios.put("/users/" + username + "/banned/" + Math.floor(Math.random() * 10000), {}, {
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
        async unbanUser(username) {
            try {
                let response = await this.$axios.get("/users/" + username + "/banned", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.ban = response.data
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
                let response = await this.$axios.delete("/users/" + username + "/banned/" + this.ban.banId, {
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
        async followUser(username) {
            try {
                let response = await this.$axios.put("/users/" + username + "/followed/" + Math.floor(Math.random() * 10000), {}, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.clear = response.data
                this.refresh()
                this.successmsg = "User" + username + "followed successfully"
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
        async unfollowUser(username) {
            try {
                let response = await this.$axios.get("/users/" + username + "/followed", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })
                this.follow = response.data
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
                let response = await this.$axios.delete("/users/" + username + "/followed/" + this.follow.followId, {
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
        async sendComment(username, photoid, comment) {
            if (comment === "") {
                this.errormsg = "Emtpy comment field."
            } else {
                try {
                    let response = await this.$axios.put("/users/" + username + "/photos/" + photoid + "/comments/" + Math.floor(Math.random() * 10000), { content: comment }, {
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
                let response = await this.$axios.get("/users/" + username + "/photos/" + photoid + "/comments", {
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
            </ul>
        </div>
    </nav>
    <div class="ms-sm-auto col-lg-10 px-4">
        <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
            <div v-if="profile.checkIfBanned == true" class="alert alert-danger w-100 mb-3" role="alert">
                <h4 class="alert-heading">{{ profile.username }} has banned you</h4>
                <hr>
            </div>

            <div v-if="profile.checkIfBanned == false" class="w-100">
                <div class="d-flex justify-content-between flex-wrap align-items-center mb-3">
                    <h1 class="h2">Profile of {{ profile.username }}</h1>
                    <div class="text-center">
                        <div class="d-flex justify-content-end mb-2">
                            <div class="px-3">
                                <p class="mb-1 h5">{{ profile.followersCount }}</p>
                                <p class="small text-muted mb-0">Followers</p>
                            </div>
                            <div class="px-3">
                                <p class="mb-1 h5">{{ profile.followingCount }}</p>
                                <p class="small text-muted mb-0">Followings</p>
                            </div>
                        </div>
                        <div class="form-group row">
                            <div class="col-md-6 mb-2">
                                <button v-if="profile.followStatus == false" class="btn btn-outline-primary w-100 text-center" @click="followUser(profile.username)">Follow</button>
                                <button v-if="profile.followStatus == true" class="btn btn-primary w-100 text-center" @click="unfollowUser(profile.username)">Unfollow</button>
                            </div>
                            <div class="col-md-6 mb-2">
                                <button v-if="profile.banStatus == false" class="btn btn-outline-danger w-100" @click="banUser(profile.username)">Ban</button>
                                <button v-if="profile.banStatus == true" class="btn btn-danger w-100" @click="unbanUser(profile.username)">Unban</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <SuccessMsg v-if="successmsg" :msg="successmsg"></SuccessMsg>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

        <div v-if="profile.checkIfBanned == false" class="row">
         <div class="photo-stream">
            <div v-for="photo in photoList.photos" :key="photo.id">
                <div class="card shadow-sm">
                    <img class="card-img-top" :src="photo.file" alt="Card image cap">
                    <div class="card-body">
                        <RouterLink :to="'/users/' + profile.username + '/view'" class="nav-link">
                            <button type="button" class="btn btn-outline-primary mb-2">{{ profile.username }}</button>
                        </RouterLink>
                        <div class="d-flex justify-content-between mb-2">
                            <p class="card-text">Likes: {{ photo.likesCount }}</p>
                            <p class="card-text">Comments: {{ photo.commentsCount }}</p>
                        </div>
                        <p class="card-text mb-2">Uploaded on: {{ photo.date }}</p>
                        <div class="input-group">
                            <input type="text" v-model="photo.comment" class="form-control" placeholder="Comment!" aria-label="Comment">
                            <button class="btn btn-primary" @click="sendComment(profile.username, photo.id, photo.comment)">Send</button>
                        </div>
                        <div class="d-flex justify-content-between align-items-center mt-2">
                            <button class="btn btn-dark" @click="openLog(profile.username, photo.id)">Comments</button>
                            <Comments id="logviewer" :log="photoComments" :token="token"></Comments>
                            <button v-if="photo.likeStatus == false" class="btn btn-primary" @click="likePhoto(profile.username, photo.id)">Like</button>
                            <button v-if="photo.likeStatus == true" class="btn btn-danger" @click="removeLike(profile.username, photo.id)">Unlike</button>
                        </div>
                    </div>
                </div>
             </div>
            </div>
           
        </div>
        
    </div>
</template>


<style scoped>
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

</style>