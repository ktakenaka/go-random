import React from "react";
import { List } from "antd";

import { TypeSample } from "constants/type";

type Props = {
  header?: string;
  samples: Array<TypeSample>;
  footer?: React.ReactNode;
};

const SampleList = ({ samples, header, footer }: Props) => {
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
