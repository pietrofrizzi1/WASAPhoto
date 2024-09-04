<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<script>
export default {
    data() {
        return {};
    },
}
</script>

<template>
    <div class="page-container">
        <!-- Navbar -->
        <header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
            <a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASAPhoto</a>
        </header>

        <div class="container-fluid flex-grow-1">
            <div class="row">
                <main>
                    <RouterView />
                </main>
            </div>
        </div>

        <!-- Footer -->
        <footer class="footer py-3">
            <div class="container">
                <span class="text-black small">&copy; 2024 WASAPhoto. All Rights Reserved. Contact the developer: frizzi.2005977@studenti.uniroma1.it</span>
            </div>
        </footer>
    </div>
</template>

<style>
/* Ensure the page takes up the full height of the viewport */
html, body {
    height: 100%;
    margin: 0;
}

/* Main container that uses Flexbox to push the footer to the bottom */
.page-container {
    display: flex;
    flex-direction: column;
    min-height: 100vh; /* Ensure the footer stays at the bottom of the viewport */
}

/* Flex-grow to ensure the main content area grows and takes up available space */
.container-fluid {
    flex-grow: 1;
}

/* Footer styling */
.footer {
    background-color: #f8f9fa; /* Optional: background color for footer */
    text-align: center;
    color: black;
	margin-top: auto;
}

/* Text styling */
.text-black {
    color: black;
}
</style>

