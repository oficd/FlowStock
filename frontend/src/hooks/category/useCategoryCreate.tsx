// import useMutation dari '@tanstack/react-query';
import { useMutation } from '@tanstack/react-query';

//import service API
import Api from '../../services/api';

// import js-cookie
import Cookies from 'js-cookie';

interface CategoryRequest {
    name: string;
}

export const useCategoryCreate = () => {

    return useMutation({
        mutationFn: async (data: CategoryRequest) => {

            const token = Cookies.get('token');

            const response = await Api.post('/api/categories', data, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            return response.data;
        }
    });
};
