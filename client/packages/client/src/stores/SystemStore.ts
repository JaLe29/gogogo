/* eslint-disable no-param-reassign */
import { create, StoreApi, UseBoundStore } from 'zustand';
import { immer } from 'zustand/middleware/immer'; 

interface Props {

}

export const useSystemStore: UseBoundStore<StoreApi<Props>> = create(
	immer<Props>(set => ({
		//
	})),
);
