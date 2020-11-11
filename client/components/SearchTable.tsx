import { Divider, Input, Table } from "antd";
import React from "react";

type Props = {
  columns: { [propname: string]: any }[];
  data: { [propname: string]: any }[] | undefined;
  searchKey: string;
};
export default function SerachTable({ columns, data, searchKey }: Props) {
  const { Search } = Input;
  return (
    <>
      <Search
        placeholder={"input search " + searchKey}
        onSearch={(value) => console.log(value)}
        enterButton
      />
      <Divider />
      <Table columns={columns} dataSource={data}></Table>
    </>
  );
}
