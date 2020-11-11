import { Button, Popconfirm, Space } from "antd";
import React from "react";
import Layout from "../../../components/Layout";
import SerachTable from "../../../components/SearchTable";
export default function Flows() {
  const columns = [
    {
      title: "Flow ID",
      dataIndex: "record_id",
      key: "id",
    },
    {
      title: "Flow Name",
      dataIndex: "name",
      key: "id",
    },
    {
      title: "Version",
      dataIndex: "version",
      key: "id",
    },
    {
      title: "Created Date",
      dataIndex: "created",
      key: "id",
    },
    {
      title: "Action",
      key: "id",
      render: (_: any, record: any) => {
        return (
          <Space size="middle">
            <Button type="text">Edit</Button>
            <Popconfirm
              title={"Sure to delete ?"}
              onConfirm={() => {
                // deleteFlow(record.id);
                console.log("deleted!", record.id);
              }}
            >
              <Button danger type="text">
                Delete
              </Button>
            </Popconfirm>
          </Space>
        );
      },
    },
  ];
  return (
    <Layout title="Home">
      <h1>Flows</h1>
      <p>这里是flow</p>
      <SerachTable columns={columns} data={[]} searchKey="" />
    </Layout>
  );
}
