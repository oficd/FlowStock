// import useMutation dari '@tanstack/react-query';
import { useMutation } from '@tanstack/react-query';

//import service API
import Api from '../../services/api';

// import js-cookie
import Cookies from 'js-cookie';

interface SupplierRequest {
    name: string;
}

export const useSupplierCreate = () => {

    return useMutation({
        mutationFn: async (data: SupplierRequest) => {

            const token = Cookies.get('token');

            const response = await Api.post('/api/suppliers', data, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            return response.data;
        }
    });
};
