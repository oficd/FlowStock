// import useMutation dari '@tanstack/react-query';
import { useMutation } from '@tanstack/react-query';

// import service API
import Api from '../../services/api';

// import js-cookie
import Cookies from 'js-cookie';

export const useCategoryDelete = () => {

    return useMutation({
        mutationFn: async (id: number) => {
            const token = Cookies.get('token');

            const response = await Api.delete(`/api/categories/${id}`, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            return response.data;
        }
    });

};
