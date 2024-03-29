/**
 * Generated by orval v6.23.0 🍺
 * Do not edit manually.
 * Bastard proxy api
 * OpenAPI spec version: 1.0.0
 */
import {
  useMutation,
  useQuery
} from '@tanstack/react-query'
import type {
  MutationFunction,
  QueryFunction,
  QueryKey,
  UseMutationOptions,
  UseQueryOptions,
  UseQueryResult
} from '@tanstack/react-query'
import axios from 'axios'
import type {
  AxiosError,
  AxiosRequestConfig,
  AxiosResponse
} from 'axios'
import type {
  DeleteApiProxyParams
} from '../model/deleteApiProxyParams'
import type {
  NewProxy
} from '../model/newProxy'
import type {
  PatchApiProxyParams
} from '../model/patchApiProxyParams'
import type {
  PatchProxy
} from '../model/patchProxy'
import type {
  Proxy
} from '../model/proxy'
import type {
  SuccessResponse
} from '../model/successResponse'



/**
 * Update a proxy
 * @summary Update a proxy
 */
export const patchApiProxy = (
    patchProxy: PatchProxy,
    params: PatchApiProxyParams, options?: AxiosRequestConfig
 ): Promise<AxiosResponse<SuccessResponse>> => {
    
    return axios.patch(
      `/api/proxy`,
      patchProxy,{
    ...options,
        params: {...params, ...options?.params},}
    );
  }



export const getPatchApiProxyMutationOptions = <TError = AxiosError<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof patchApiProxy>>, TError,{data: PatchProxy;params: PatchApiProxyParams}, TContext>, axios?: AxiosRequestConfig}
): UseMutationOptions<Awaited<ReturnType<typeof patchApiProxy>>, TError,{data: PatchProxy;params: PatchApiProxyParams}, TContext> => {
 const {mutation: mutationOptions, axios: axiosOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof patchApiProxy>>, {data: PatchProxy;params: PatchApiProxyParams}> = (props) => {
          const {data,params} = props ?? {};

          return  patchApiProxy(data,params,axiosOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type PatchApiProxyMutationResult = NonNullable<Awaited<ReturnType<typeof patchApiProxy>>>
    export type PatchApiProxyMutationBody = PatchProxy
    export type PatchApiProxyMutationError = AxiosError<unknown>

    /**
 * @summary Update a proxy
 */
export const usePatchApiProxy = <TError = AxiosError<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof patchApiProxy>>, TError,{data: PatchProxy;params: PatchApiProxyParams}, TContext>, axios?: AxiosRequestConfig}
) => {

      const mutationOptions = getPatchApiProxyMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * Create a new proxy
 * @summary Create a new proxy
 */
export const postApiProxy = (
    newProxy: NewProxy, options?: AxiosRequestConfig
 ): Promise<AxiosResponse<SuccessResponse>> => {
    
    return axios.post(
      `/api/proxy`,
      newProxy,options
    );
  }



export const getPostApiProxyMutationOptions = <TError = AxiosError<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof postApiProxy>>, TError,{data: NewProxy}, TContext>, axios?: AxiosRequestConfig}
): UseMutationOptions<Awaited<ReturnType<typeof postApiProxy>>, TError,{data: NewProxy}, TContext> => {
 const {mutation: mutationOptions, axios: axiosOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof postApiProxy>>, {data: NewProxy}> = (props) => {
          const {data} = props ?? {};

          return  postApiProxy(data,axiosOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type PostApiProxyMutationResult = NonNullable<Awaited<ReturnType<typeof postApiProxy>>>
    export type PostApiProxyMutationBody = NewProxy
    export type PostApiProxyMutationError = AxiosError<unknown>

    /**
 * @summary Create a new proxy
 */
export const usePostApiProxy = <TError = AxiosError<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof postApiProxy>>, TError,{data: NewProxy}, TContext>, axios?: AxiosRequestConfig}
) => {

      const mutationOptions = getPostApiProxyMutationOptions(options);

      return useMutation(mutationOptions);
    }
    /**
 * Get all proxies
 * @summary Get all proxies
 */
export const getApiProxy = (
     options?: AxiosRequestConfig
 ): Promise<AxiosResponse<Proxy[]>> => {
    
    return axios.get(
      `/api/proxy`,options
    );
  }


export const getGetApiProxyQueryKey = () => {
    return [`/api/proxy`] as const;
    }

    
export const getGetApiProxyQueryOptions = <TData = Awaited<ReturnType<typeof getApiProxy>>, TError = AxiosError<unknown>>( options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getApiProxy>>, TError, TData>>, axios?: AxiosRequestConfig}
) => {

const {query: queryOptions, axios: axiosOptions} = options ?? {};

  const queryKey =  queryOptions?.queryKey ?? getGetApiProxyQueryKey();

  

    const queryFn: QueryFunction<Awaited<ReturnType<typeof getApiProxy>>> = ({ signal }) => getApiProxy({ signal, ...axiosOptions });

      

      

   return  { queryKey, queryFn, ...queryOptions} as UseQueryOptions<Awaited<ReturnType<typeof getApiProxy>>, TError, TData> & { queryKey: QueryKey }
}

export type GetApiProxyQueryResult = NonNullable<Awaited<ReturnType<typeof getApiProxy>>>
export type GetApiProxyQueryError = AxiosError<unknown>

/**
 * @summary Get all proxies
 */
export const useGetApiProxy = <TData = Awaited<ReturnType<typeof getApiProxy>>, TError = AxiosError<unknown>>(
  options?: { query?:Partial<UseQueryOptions<Awaited<ReturnType<typeof getApiProxy>>, TError, TData>>, axios?: AxiosRequestConfig}

  ):  UseQueryResult<TData, TError> & { queryKey: QueryKey } => {

  const queryOptions = getGetApiProxyQueryOptions(options)

  const query = useQuery(queryOptions) as  UseQueryResult<TData, TError> & { queryKey: QueryKey };

  query.queryKey = queryOptions.queryKey ;

  return query;
}



/**
 * Delete a proxy
 * @summary Delete a proxy
 */
export const deleteApiProxy = (
    params: DeleteApiProxyParams, options?: AxiosRequestConfig
 ): Promise<AxiosResponse<SuccessResponse>> => {
    
    return axios.delete(
      `/api/proxy`,{
    ...options,
        params: {...params, ...options?.params},}
    );
  }



export const getDeleteApiProxyMutationOptions = <TError = AxiosError<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteApiProxy>>, TError,{params: DeleteApiProxyParams}, TContext>, axios?: AxiosRequestConfig}
): UseMutationOptions<Awaited<ReturnType<typeof deleteApiProxy>>, TError,{params: DeleteApiProxyParams}, TContext> => {
 const {mutation: mutationOptions, axios: axiosOptions} = options ?? {};

      


      const mutationFn: MutationFunction<Awaited<ReturnType<typeof deleteApiProxy>>, {params: DeleteApiProxyParams}> = (props) => {
          const {params} = props ?? {};

          return  deleteApiProxy(params,axiosOptions)
        }

        


   return  { mutationFn, ...mutationOptions }}

    export type DeleteApiProxyMutationResult = NonNullable<Awaited<ReturnType<typeof deleteApiProxy>>>
    
    export type DeleteApiProxyMutationError = AxiosError<unknown>

    /**
 * @summary Delete a proxy
 */
export const useDeleteApiProxy = <TError = AxiosError<unknown>,
    TContext = unknown>(options?: { mutation?:UseMutationOptions<Awaited<ReturnType<typeof deleteApiProxy>>, TError,{params: DeleteApiProxyParams}, TContext>, axios?: AxiosRequestConfig}
) => {

      const mutationOptions = getDeleteApiProxyMutationOptions(options);

      return useMutation(mutationOptions);
    }
    