import { Button, Col, Modal, Row, Space, Table, Typography } from 'antd';
import styled from 'styled-components';
import { Link } from 'react-router-dom';
import { useQuery } from '@apollo/client';
import Page from './Page';
import useAxios from 'axios-hooks';
import axios from 'axios';
import SimpleForm from '../components/SimpleForm';
import { useState } from 'react';
import { BASE_URL } from '../const/client';


const ProxyPage: React.FC = () => {
    const [isModalOpen, setIsModalOpen] = useState(false);

    const showModal = () => {
        setIsModalOpen(true);
    };

    const handleOk = () => {
        setIsModalOpen(false);
    };

    const handleCancel = () => {
        setIsModalOpen(false);
    };

    const [{ data, loading, error }, refetch] = useAxios(
        BASE_URL + '/api/proxy',
    )

    console.log(data, loading, error)

    const onDelete = async (id: string) => {
        await axios.delete(BASE_URL + `/api/proxy?id=${id}`)
 
        refetch()
    }

    const onCreate = async (values: { source: string, target: string }) => {
        await axios.post(BASE_URL + `/api/proxy`, JSON.stringify(values))

        refetch()
        handleCancel();
    }

    const columns = [
        {
            title: 'Id',
            dataIndex: 'id',
            key: 'id',
        }, {
            title: 'Source',
            dataIndex: 'source',
            key: 'source',
        },
        {
            title: 'Target',
            dataIndex: 'target',
            key: 'target',
        },
        {
            title: 'Actions',
            key: 'actions',
            render: (text: any, record: any) => (
                <Space size="middle">
                    <Button type="primary" danger onClick={() => onDelete(record.id)}>Delete</Button>
                </Space>
            ),
        },

    ]

    return (
        // <BasicTransition>
        <Page>
            <Button type="primary" onClick={showModal}>
                Open Modal
            </Button>
            <Modal footer={null} title="Basic Modal" open={isModalOpen} onCancel={handleCancel}  >
                <SimpleForm submitButtonTitle='Create' fields={[{ key: "source", type: "string", label: "Source", rules: [{ required: true }] }, { key: "target", type: "string", label: "Target", rules: [{ required: true }] }]} onFinish={onCreate} />
            </Modal>
            <Table dataSource={data} columns={columns} loading={loading} />
        </Page>
        // </BasicTransition >
    );
};

export default ProxyPage;
