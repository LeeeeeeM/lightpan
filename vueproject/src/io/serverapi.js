 
import { axiosfetch } from '../uitl/fetch';

export const login = data => axiosfetch('/login', data, 'POST')