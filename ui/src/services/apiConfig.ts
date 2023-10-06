import axios from 'axios';

export const mlServerApi = axios.create({
    baseURL: process.env.NEXT_PUBLIC_ML_SERVER_BASE_URL || 'http://xpuls-ml-server',
});