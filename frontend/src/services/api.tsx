//import axios
import axios from 'axios';
import Cookies from 'js-cookie';

// handle unauthorized globally
function handleUnauthorized() {
    try {
        Cookies.remove('token');
        Cookies.remove('user');
    } catch {}
    try { localStorage.setItem('logout', Date.now().toString()); } catch {}
    // redirect to login
    window.location.href = '/login';
}

const Api = axios.create({
    //set default endpoint API
    baseURL: 'http://localhost:3000'
})

// attach token automatically
Api.interceptors.request.use((config) => {
    const token = Cookies.get('token');
    if (token && config.headers) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

// handle 401 globally
Api.interceptors.response.use(
    (res) => res,
    (error) => {
        const status = error?.response?.status;
        if (status === 401) {
            handleUnauthorized();
        }
        return Promise.reject(error);
    }
);

export default Api
