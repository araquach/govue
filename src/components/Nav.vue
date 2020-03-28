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
                <b-navbar-item tag="router-link" to="/">
                    Home
                </b-navbar-item>
                <b-navbar-item tag="router-link" to="/about">
                    About
                </b-navbar-item>
            </b-navbar-dropdown>
        </template>
        <template slot="end">
            <b-navbar-item tag="div">
                <div class="buttons">
                    <button tag="router-link" to="/" class="button is-primary">
                        Login
                    </button>
                </div>
            </b-navbar-item>
        </template>
    </b-navbar>
</template>

<script>

</script>