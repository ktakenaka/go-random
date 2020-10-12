import React from "react";
import { List, Popconfirm } from "antd";

import { TypeSample } from "constants/type";

type Props = {
  header?: string;
  samples: Array<TypeSample>;
  footer?: React.ReactNode;
  onDelete: (id: number) => void;
};

const SampleList = ({ samples, header, footer, onDelete }: Props) => {
  return (
    <List
      itemLayout="horizontal"
      header={header}
      dataSource={samples}
      renderItem={(item) => (
        <List.Item
          actions={[
            <a key={item.id} href={`/samples/${item.id}/edit`}>
              edit
            </a>,
            <Popconfirm
              id={String(item.id)}
              title="Are you sure delete this sample?"
              onConfirm={() => onDelete(item.id!)}
              okText="Yes"
              cancelText="No"
            >
              <a href="#">delete</a>
            </Popconfirm>,
          ]}
        >
          <List.Item.Meta title={item.id} description={item.title} />
          <div>{item.content}</div>
        </List.Item>
      )}
      footer={footer}
    />
  );
};

export default SampleList;
