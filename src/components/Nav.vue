<template>
    <b-navbar class="is-transparent">
        <template slot="start">
            <b-navbar-item tag="router-link" to="/">
                Home
            </b-navbar-item>
            <b-navbar-item tag="router-link" to="/about">
                About
            </b-navbar-item>
            <b-navbar-dropdown label="More">
                <b-navbar-item tag="router-link" to="/link1">
                    Link 1
                </b-navbar-item>
                <b-navbar-item tag="router-link" to="/link2">
                    Link 2
                </b-navbar-item>
            </b-navbar-dropdown>
        </template>
        <template slot="end">
            <b-navbar-item tag="div">
                <div class="buttons">
                    <router-link class="button is-warning" to="/register" >
                        Register
                    </router-link>
                    <router-link class="button is-primary" to="/login">
                        Login
                    </router-link>
                </div>
            </b-navbar-item>
        </template>
    </b-navbar>
</template>

<script>

</script>