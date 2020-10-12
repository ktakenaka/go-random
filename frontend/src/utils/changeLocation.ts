import history from "browserHistory";

export const moveLocation = (location: string) => {
  history.push(location);
};
