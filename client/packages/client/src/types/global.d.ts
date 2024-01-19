export {};
type _AnyReactElement = ReactFragment | ReactPortal | boolean | string | null | undefined | React.ReactElement;

declare global {
	type AnyReactElement = _AnyReactElement | _AnyReactElement[];
}
