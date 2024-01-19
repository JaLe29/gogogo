/* eslint-disable no-console */
/* eslint-disable react/no-array-index-key */
import { ApolloClient, DefaultOptions, from, HttpLink, InMemoryCache } from '@apollo/client';
import { onError } from '@apollo/client/link/error';
import { useMemo } from 'react';
import { BASE_API } from './app';

const defaultOptions: DefaultOptions = {
	watchQuery: {
		fetchPolicy: 'no-cache',
		errorPolicy: 'ignore',
	},
	query: {
		fetchPolicy: 'no-cache',
		errorPolicy: 'all',
	},
};

export const useApolloClient = (token?: string): ApolloClient<any> => {
	const client = useMemo(
		() =>
			new ApolloClient({
				defaultOptions,
				cache: new InMemoryCache(),
				link: from([
					onError(e => {
						console.error(e);
						const msgs = e.graphQLErrors?.map(gqlE => gqlE.message);
						if (msgs?.length) {
							// notification.open({
							// 	type: 'error',
							// 	description: (
							// 		<>
							// 			{msgs.map((msg, i) => (
							// 				<p key={i}>{msg}</p>
							// 			))}
							// 		</>
							// 	),
							// 	message: 'Nastala chyba!',
							// 	closeIcon: false,
							// });
							// return
						}
						// notification.open({
						// 	type: 'error',
						// 	description: (
						// 		<>
						// 			<p>Omlouváme se, nastala chyba</p>
						// 		</>
						// 	),
						// 	message: 'Nastala chyba!',
						// 	closeIcon: false,
						// });
					}),
					new HttpLink({
						uri: `${BASE_API}/graphql`,
						fetch: (url: any, init: any): any =>
							fetch(url, {
								...init,
								headers: {
									...init.headers,
									authorization: token || init.headers.authorization,
								},
							}).then(response => response),
					}),
				]),
			}),
		[token],
	);
	return client;
};
