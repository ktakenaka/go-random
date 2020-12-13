import React from "react";
import { List, Button, Popconfirm } from "antd";

import { TypeSample } from "constants/type";
import { moveLocation } from "utils/changeLocation";

type Props = {
  header?: string;
  samples: Array<TypeSample>;
  footer?: React.ReactNode;
  onDelete: (id: string) => void;
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
            <Button
              type="link"
              size="small"
              key={item.id + "edit"}
              onClick={() => moveLocation(`/samples/${item.id}/edit`)}
            >
              edit
            </Button>,
            <Popconfirm
              key={item.id + "delete"}
              title="Are you sure delete this sample?"
              onConfirm={() => onDelete(String(item.id))}
              okText="Yes"
              cancelText="No"
            >
              <Button type="link" size="small">
                delete
              </Button>
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
