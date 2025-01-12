import {ComponentStory, ComponentMeta} from '@storybook/react';

import CreateTestSteps from '../CreateSteps';

export default {
  title: 'Create Test Steps',
  component: CreateTestSteps,
} as ComponentMeta<typeof CreateTestSteps>;

const Template: ComponentStory<typeof CreateTestSteps> = args => <CreateTestSteps {...args} />;

export const Default = Template.bind({});
Default.args = {
  selectedKey: '2',
  componentFactory(step) {
    return <div>{step}</div>;
  },
  stepList: [
    {
      id: '1',
      title: 'Create Test',
      name: 'Create Test',
      status: 'complete',
      component: 'SelectPlugin',
    },
    {
      id: '2',
      title: 'Basic Details',
      name: 'Create Test',
      status: 'selected',
      component: 'BasicDetails',
    },
    {
      id: '3',
      title: 'Request Details',
      name: 'Create Test',
      status: 'pending',
      component: 'RequestDetails',
    },
  ],
};
