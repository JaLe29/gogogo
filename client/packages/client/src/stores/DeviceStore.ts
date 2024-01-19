/* eslint-disable no-param-reassign */
import { create, StoreApi, UseBoundStore } from 'zustand';
import { immer } from 'zustand/middleware/immer';

interface Props {
	isNavigationOpened: boolean;
	setNavigationOpened: (opened: boolean) => void;
}

export const useDeviceStore: UseBoundStore<StoreApi<Props>> = create(
	immer<Props>(set => ({
		isNavigationOpened: true,
		setNavigationOpened: opened => set({ isNavigationOpened: opened }),
	})),
);
