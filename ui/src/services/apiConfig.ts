import axios from 'axios';

export const mlServerApi = axios.create({
    baseURL: `/api`,
});
