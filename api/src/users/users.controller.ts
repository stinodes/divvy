import {
  Controller,
  Get,
  Post,
  Body,
  Param,
  NotFoundException,
} from '@nestjs/common';

import { UsersService } from './users.service';
import { CreateUserDto } from './dto/create-user.dto';
import { AddFriendDto } from './dto/add-friend.dto';

@Controller('users')
export class UsersController {
  constructor(private readonly usersService: UsersService) {}

  @Post()
  async create(@Body() createUserDto: CreateUserDto) {
    const auth0User = await this.usersService.findOne(createUserDto.user_id);
    if (!auth0User) return;
    await this.usersService.create(createUserDto);
  }

  @Get()
  async findAll() {
    const data = await this.usersService.findAll();
    return data;
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.usersService.findOne(id);
  }

  @Get(':id/friends')
  findFriends(@Param('id') id: string) {
    return this.usersService.findFriends(id);
  }

  @Post(':id/friends')
  async addFriend(@Param('id') id: string, @Body() body: AddFriendDto) {
    const friend = await this.usersService.findOne(body.user_id);
    if (!friend) throw new NotFoundException();
    await this.usersService.addFriend(id, body.user_id);
  }
}
