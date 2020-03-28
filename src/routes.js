import Home from './components/home'
import About from './components/About'
import Login from './components/auth/Login'
import Register from  './components/auth/Register'

export const routes = [
    { path: '', component: Home },
    { path: '/login', component: Login },
    { path: '/about', component: About },
    { path: '/register', component: Register }
]

