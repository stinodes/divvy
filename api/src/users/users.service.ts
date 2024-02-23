import { Injectable, NotFoundException } from '@nestjs/common';
import { CreateUserDto } from './dto/create-user.dto';
import { InjectRepository } from '@nestjs/typeorm';
import { User } from './entities/user.entity';
import { Repository } from 'typeorm';
import { Auth0Service } from 'src/auth/auth0.service';

@Injectable()
export class UsersService {
  constructor(
    @InjectRepository(User)
    private usersRepository: Repository<User>,
    private auth0Service: Auth0Service,
  ) {}

  async create(createUserDto: CreateUserDto) {
    const user = new User();
    user.auth0_id = createUserDto.user_id;
    user.friends = [];
    await this.usersRepository.save(user);
  }

  findAll() {
    return this.auth0Service.findAllAuth0Users({});
  }
  findOne(id: string) {
    return this.auth0Service.findOneAuth0User(id);
  }

  async findFriends(auth0_id: string) {
    const user = await this.usersRepository.findOne({
      where: { auth0_id },
      relations: { friends: true },
    });
    const results = await Promise.all(
      (user?.friends || []).map((user) => {
        const id = user.auth0_id;
        return this.auth0Service.findOneAuth0User(id);
      }),
    );
    return results.filter(Boolean);
  }

  async addFriend(id: string, friendId: string) {
    const user = await this.usersRepository.findOne({
      where: { auth0_id: id },
      relations: { friends: true },
    });
    const friend = await this.usersRepository.findOne({
      where: { auth0_id: friendId },
    });
    if (!user || !friend) throw new NotFoundException();
    if (!user.friends.some((v) => v.auth0_id === friendId)) {
      user.friends.push(friend);
      await this.usersRepository.save(user);
    }
  }
}
