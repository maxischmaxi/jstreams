export type Team = {
  bans: Ban[];
  objectives: Objectives;
  teamId: number;
  win: boolean;
};

export type Ban = {
  championId: number;
  pickTurn: number;
};

export type Objectives = {
  baron: Baron;
  champion: Champion;
  dragon: Dragon;
  horde: Horde;
  inhibitor: Inhibitor;
  riftHerald: RiftHerald;
  tower: Tower;
};

export type Baron = {
  first: boolean;
  kills: number;
};

export type Champion = {
  first: boolean;
  kills: number;
};

export type Dragon = {
  first: boolean;
  kills: number;
};

export type Horde = {
  first: boolean;
  kills: number;
};

export type Inhibitor = {
  first: boolean;
  kills: number;
};

export type RiftHerald = {
  first: boolean;
  kills: number;
};

export type Tower = {
  first: boolean;
  kills: number;
};
