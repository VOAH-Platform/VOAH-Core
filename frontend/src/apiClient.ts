import axios, { AxiosError } from 'axios';

export const API_HOST = !import.meta.env.DEV
  ? window.location.href
  : 'https://test-voah.zirr.al';

export const apiClient = axios.create({
  baseURL: API_HOST,
});
apiClient.interceptors.response.use(
  (res) => res,
  (err: AxiosError) => {
    const response = err.response;
    if (!response) return Promise.reject(err);
    if (response.status === 401 && response.data === 'Invalid or expired JWT') {
      const userData = JSON.parse(localStorage.getItem('user')!) as {
        email: string;
        id: string;
        isLogin: boolean;
        accessToken: string;
        refreshToken: string;
      };
      // eslint-disable-next-line promise/no-promise-in-callback
      axios
        .post<{
          'access-token': string;
          message: string;
          exp: number;
        }>(`${API_HOST}/api/auth/refresh`, {
          'user-id': userData.id,
          'refresh-token': userData.refreshToken,
        })
        .then((res) => {
          localStorage.setItem(
            'user',
            JSON.stringify({
              email: userData.email,
              id: userData.id,
              isLogin: true,
              accessToken: res.data['access-token'],
              refreshToken: userData.refreshToken,
            }),
          );
          window.location.reload();
          return;
        })
        .catch((err) => {
          console.log(err);
        });
    }
    return Promise.reject(err);
  },
);
