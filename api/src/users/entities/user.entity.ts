import { Entity, JoinTable, ManyToMany, PrimaryColumn } from 'typeorm';

@Entity()
export class User {
  @PrimaryColumn()
  auth0_id: string;

  @ManyToMany(() => User)
  @JoinTable()
  friends: User[];
}
