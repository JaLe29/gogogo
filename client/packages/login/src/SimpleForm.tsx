import { Alert, Button, Form, Input, Select, Checkbox, DatePicker } from 'antd';
import { Rule } from 'antd/es/form';
import { Store } from 'antd/es/form/interface';
import { Dayjs } from 'dayjs';
import { useState } from 'react';
import styled from 'styled-components';

const ItemWrapper = styled.div`
	margin-inline-end: 0px;
	padding-top: 10px;
	padding-bottom: 10px;
	width: 100%;
	margin-bottom: 0 !important;
`;

const StyledDatePicker = styled(DatePicker)`
	width: 100%;
`;

const FormItemStyled = styled(Form.Item)`
	margin-bottom: 0 !important;

	&:first-child {
		width: 100%;
	}

	label {
		width: 100%;
	}

	.ant-form-item-required::before {
		display: none !important;
	}

	.ant-row {
		display: flex;
		flex-direction: column;
	}
`;

const CheckboxStyled = styled(Checkbox)`
	padding: 10px;
	width: 100%;
`;

const FormItemButtonStyled = styled(Form.Item)`
	padding-top: 10px;
	width: 100%;

	button {
		width: 100%;
	}
`;

type BaseField = {
    rules?: Rule[];
    footer?: any;
};

type FieldString = {
    label: any;
    key: string;
    required?: boolean;
    placeholder?: string;
    type?: 'string' | 'password';
    /**
     * Převede hodnotu na požadovaný formát
     * @param value
     * @returns
     */
    normalize?: (value: any) => any;
} & BaseField;

type FieldDayPicker = {
    label: any;
    key: string;
    required?: boolean;
    placeholder?: string;
    maxDate?: Dayjs;
    minDate?: Dayjs;
} & BaseField;

type FieldCheckBoxGroup = {
    label: any;
    key: string;
    required?: boolean;
    type: 'checkbox-group';
    values: { value: string; key: string; disabled?: boolean }[];
} & BaseField;

type FieldSelect = {
    label: any;
    key: string;
    required?: boolean;
    placeholder?: string;
    type: 'select';

    mode?: 'multiple' | 'tags';
    /**
     * Převede hodnotu na požadovaný formát
     * @param value
     * @returns
     */
    normalize?: (value: any) => any;
} & BaseField &
    (
        | { groups?: boolean; values: { value: string; key: string; group: string; disabled?: boolean }[] }
        | { values: { value: string; key: string; disabled?: boolean }[] }
    );

type FieldCheckBox = {
    key: string;
    required?: boolean;
    type: 'checkbox';
    node: any;
} & BaseField;

export type Field = FieldDayPicker | FieldString | FieldSelect | FieldCheckBox | FieldCheckBoxGroup;

const isFieldSelect = (field: Field): field is FieldSelect => (field as any).type === 'select';
const isFieldCheckBox = (field: Field): field is FieldCheckBox => (field as any).type === 'checkbox';

interface Props {
    hideLabels?: boolean;
    submitButtonTitle?: string;
    fields: Field[];
    onFinish: (values: any) => void;
    onChange?: (values: any) => void;
    onFinishFailed?: (errorInfo: any) => void;
    initialValues?: Store;
    loading?: boolean;
    customButton?: any;
    disableSubmitButtonIfInitialEqual?: boolean;
    rules?: Record<string, (changed: unknown) => { key: string; value: unknown }[]>;
}

const SimpleForm: React.FC<Props> = ({
    onFinish,
    fields,
    submitButtonTitle,
    onFinishFailed,
    loading,
    initialValues,
    customButton,
    onChange,
    disableSubmitButtonIfInitialEqual,
    hideLabels = false,
    rules = {},
}) => {
    const [disabled, setDisabled] = useState(disableSubmitButtonIfInitialEqual ?? false);
    const [form] = Form.useForm();

    const resolveField = (field: any): any => {
        if (field.type === 'string' || !field.type) {
            return <Input placeholder={field.placeholder} size="large" />;
        }

        if (isFieldSelect(field)) {
            if ((field as any).groups) {
                const groups = (field as any).values as {
                    value: string;
                    key: string;
                    group: string;
                    disabled?: boolean;
                }[];
                const options = groups.reduce((acc: any, v) => {
                    const group = acc.find((g: { label: string }) => g.label === v.group);
                    if (group) {
                        group.options.push({ label: v.value, value: v.key, disabled: v.disabled });
                    } else {
                        acc.push({ label: v.group, options: [{ label: v.value, value: v.key, disabled: v.disabled }] });
                    }
                    return acc;
                }, []);
                // console.log(field.placeholder);
                return (
                    <Select
                        style={{ width: '100%' }}
                        size="large"
                        placeholder={field.placeholder}
                        mode={field.mode}
                        options={options}
                        showSearch
                        filterOption={(
                            input: string,
                            option?: {
                                label: string;
                                value: string;
                                children: any;
                                options: { label: string; value: string; children: any }[];
                            },
                        ) => {
                            if (option?.options) {
                                return false;
                            }
                            const label = option?.label;
                            if (typeof label === 'string') {
                                return (label ?? '').toLowerCase().includes(input.toLowerCase());
                            }
                            return false;
                        }}
                    />
                );
            }

            return (
                <Select
                    style={{ width: '100%' }}
                    size="large"
                    placeholder={field.placeholder}
                    mode={field.mode}
                    showSearch
                    filterOption={(
                        input: string,
                        option?: { label: string; value: string; children: any },
                    ) => {
                        const children = option?.children;
                        if (typeof children === 'string') {
                            return (children ?? '').toLowerCase().includes(input.toLowerCase());
                        }
                        return false;
                    }}
                >
                    {field.values.map(value => (
                        <Select.Option key={value.key} value={value.key} disabled={value.disabled}>
                            {value.value}
                        </Select.Option>
                    ))}
                </Select>
            );
        }

        if (field.type === 'password') {
            return <Input.Password placeholder={field.placeholder} size="large" />;
        }

        if (field.type === 'checkbox-group') {
            return (
                <Checkbox.Group style={{ width: '100%' }}>
                    {field.values.map((value: any) => (
                        <CheckboxStyled value={value.key} key={value.key} disabled={value.disabled}>
                            {value.value}
                        </CheckboxStyled>
                    ))}
                </Checkbox.Group>
            );
        }

        if (isFieldCheckBox(field)) {
            return <CheckboxStyled>{field.node}</CheckboxStyled>;
        }

        return <Alert description={`Field ${field.key} type is not specified`} />;
    };

    const onHandleFinish = (v: any): void => {
        onFinish(v);
    };

    return (
        <Form
            layout="inline"
            onFinish={onHandleFinish}
            onFinishFailed={onFinishFailed}
            autoComplete="off"
            initialValues={initialValues}
            disabled={loading}
            validateTrigger="onSubmit"
            form={form}
            onValuesChange={(changed, allValues) => {
                const chengedKeys = Object.keys(changed);
                const rulesKeys = Object.keys(rules);
                const keys = chengedKeys.filter(key => rulesKeys.includes(key));

                const toSetValues: { name: string; value: any }[] = [];
                keys.forEach(key => {
                    const rule = rules[key];
                    if (typeof rule === 'function') {
                        const value = changed[key];
                        const result = rule(value);

                        result.forEach((v: { key: string; value: unknown }) => {
                            toSetValues.push({ name: v.key, value: v.value });
                        });
                    }
                });
                let allowChange = true;
                if (toSetValues.length > 0) {
                    form.setFields(toSetValues);
                    onChange?.(form.getFieldsValue());
                    allowChange = false;
                }

                if (disableSubmitButtonIfInitialEqual) {
                    setDisabled(JSON.stringify(allValues) === JSON.stringify(initialValues));
                }
                if (allowChange) {
                    onChange?.(allValues);
                }
            }}
        >
            {fields.map((field: any) => {
                let label;
                if (!hideLabels) {
                    label = field.label
                        ? `${field.label} ${field.required === false ? ' (nepovinné)' : ''}`
                        : undefined;
                }
                return (
                    <ItemWrapper key={field.key}>
                        <FormItemStyled
                            key={field.key}
                            label={label}
                            colon={false}
                            name={field.key}
                            rules={field.rules}
                            normalize={field.normalize}
                        >
                            {resolveField(field)}
                        </FormItemStyled>
                        {field.footer && field.footer}
                    </ItemWrapper>
                );
            })}
            {submitButtonTitle && (
                <FormItemButtonStyled>
                    <Button
                        type="primary"
                        htmlType="submit"
                        size="large"
                        loading={loading}
                        disabled={loading || disabled}
                    >
                        {submitButtonTitle}
                    </Button>
                </FormItemButtonStyled>
            )}
            {customButton && customButton}
        </Form>
    );
};

export default SimpleForm;