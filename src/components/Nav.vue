<template>
    <b-navbar class="is-primary is-spaced">
        <template slot="brand">
            <b-navbar-item tag="router-link" :to="{ path: '/' }">
                <h1 class="title is-4 has-text-white">GoVue</h1>
                <img src="/dist/img/icons/home.svg" width="40">
            </b-navbar-item>
        </template>
        <template slot="start">
            <b-navbar-item tag="router-link" to="/">
                Home
            </b-navbar-item>
            <b-navbar-item tag="router-link" to="/test">
                Test
            </b-navbar-item>
            <b-navbar-dropdown label="More">
                <b-navbar-item tag="router-link" to="/">
                    Home
                </b-navbar-item>
                <b-navbar-item tag="router-link" to="/test">
                    Test
                </b-navbar-item>
            </b-navbar-dropdown>
        </template>
        <template slot="end">
            <b-navbar-item tag="div">
                <div class="buttons">
                    <button tag="router-link" to="/" class="button is-warning">
                        Button
                    </button>
                </div>
            </b-navbar-item>
        </template>
    </b-navbar>
</template>

<script>

</script>