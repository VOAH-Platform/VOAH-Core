import axios, { AxiosError } from 'axios';

export const API_HOST = !import.meta.env.DEV
  ? window.location.origin
  : 'https://test-voah.zirr.al';

export const apiClient = axios.create({
  baseURL: API_HOST,
});
apiClient.interceptors.response.use(
  (res) => res,
  async (err: AxiosError) => {
    const response = err.response;
    if (!response) return Promise.reject(err);
    if (response.status === 401 && response.data === 'Invalid or expired JWT') {
      const userData = JSON.parse(localStorage.getItem('voah__user')!) as {
        email: string;
        id: string;
        isLogin: boolean;
        accessToken: string;
        refreshToken: string;
      };
      // eslint-disable-next-line promise/no-promise-in-callback
      try {
        const res = await axios.post<{
          'access-token': string;
          message: string;
          exp: number;
        }>(`${API_HOST}/api/auth/refresh`, {
          'user-id': userData.id,
          'refresh-token': userData.refreshToken,
        });

        localStorage.setItem(
          'voah__user',
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
      } catch (err) {
        alert(err);
        localStorage.removeItem('voah__user');
        window.location.reload();
        throw err;
      }
    }
    return err;
  },
);
