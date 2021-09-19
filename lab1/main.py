from key import check_license
import sys
from tkinter import *
from tkinter import messagebox as mb
import pygame
import random
from math import pi

if not check_license("license.key"):
    mb.showerror("error", "Access denided! U don't have a license")
    sys.exit()
else:
    pygame.init()

    black = [0, 0, 0]
    white = [255, 255, 255]
    red = [255, 0, 0]
    green = [0,255,127]
    blue = [173,216,230]
    k = -100
    i = 1
    size = [1000, 500]
    screen = pygame.display.set_mode(size)

    # Устанавливаем заголовок окна
    pygame.display.set_caption("Человечик")

    def line1(k):
        pygame.draw.rect(screen, blue, [0, 0, 1000, 250])
        pygame.draw.rect(screen, green, [0, 250, 1000, 100])
        pygame.draw.rect(screen, green, [0, 450, 1000, 50])
        pygame.draw.line(screen, black, [100 + k, 200], [100 + k, 300], 5)
        pygame.draw.line(screen, black, [50 + k, 400], [100 + k, 300], 5)
        pygame.draw.line(screen, black, [100 + k, 300], [150 + k, 400], 5)
        pygame.draw.circle(screen, black, [100 + k, 160], 40, 5)
        pygame.draw.circle(screen, black, [120 + k, 150], 5)
        pygame.draw.arc(screen,red,[115+k,160,15,30],pi ,3*pi/2,5)
        pygame.draw.line(screen, black, [100 + k, 210], [130 + k, 250], 5)
        pygame.draw.line(screen, black, [130 + k, 250], [160 + k, 270], 5)
        pygame.draw.line(screen, black, [100 + k, 210], [70 + k, 250], 5)
        pygame.draw.line(screen, black, [70 + k, 250], [55 + k, 280], 5)
        pygame.draw.line(screen, black, [0 , 350], [1000, 350], 5)  # road
        pygame.draw.line(screen, black, [0 , 450], [1000 , 450], 5)  # road
        k += 10
        return k

    def line2(k):
        pygame.draw.rect(screen, blue, [0, 0, 1000, 250])
        pygame.draw.rect(screen, green, [0, 250, 1000, 100])
        pygame.draw.rect(screen, green, [0, 450, 1000, 50])
        pygame.draw.line(screen, black, [100 + k, 200], [100 + k, 300], 5)  # body
        pygame.draw.circle(screen, black, [100 + k, 160], 40, 5)  # head
        pygame.draw.circle(screen, black, [120 + k, 150], 5)
        pygame.draw.arc(screen,red,[115+k,160,15,30],pi ,3*pi/2,5)
        pygame.draw.line(screen, black, [100 + k, 300], [100 + k, 350], 5)  # left leg part 1
        pygame.draw.line(screen, black, [100 + k, 350], [50 + k, 380], 5)  # left leg part 2
        pygame.draw.line(screen, black, [100 + k, 300], [125 + k, 400], 5)  # right leg
        pygame.draw.line(screen, black, [100 + k, 210], [120 + k, 250], 5)  # right  hand 1
        pygame.draw.line(screen, black, [120 + k, 250], [150 + k, 280], 5)  # right  hand 1
        pygame.draw.line(screen, black, [100 + k, 210], [85 + k, 250], 5)  # left  hand 1
        pygame.draw.line(screen, black, [85 + k, 250], [85 + k, 280], 5)  # left  hand 1
        pygame.draw.line(screen, black, [0, 350], [1000 , 350], 5)  # road
        pygame.draw.line(screen, black, [0 , 450], [1000, 450], 5)  # road
        k+=10
        return k

    def line3(k):
        pygame.draw.rect(screen, blue, [0, 0, 1000, 250])
        pygame.draw.rect(screen, green, [0, 250, 1000, 100])
        pygame.draw.rect(screen, green, [0, 450, 1000, 50])
        pygame.draw.line(screen, black, [100 + k, 200], [100 + k, 300], 5)  # body
        pygame.draw.circle(screen, black, [100 + k, 160], 40, 5)  # head
        pygame.draw.circle(screen, black, [120 + k, 150], 5)
        pygame.draw.arc(screen,red,[115+k,160,15,30],pi ,3*pi/2,5)
        pygame.draw.line(screen, black, [100 + k, 300], [100 + k, 400], 5)  # left leg part
        pygame.draw.line(screen, black, [100 + k, 300], [125 + k, 350], 5)  # right leg part 1
        pygame.draw.line(screen, black, [125 + k, 350], [110 + k, 380], 5)  # right leg part 2
        pygame.draw.line(screen, black, [100 + k, 210], [120 + k, 280], 5)  # right  hand 1
        pygame.draw.line(screen, black, [100 + k, 210], [80 + k, 280], 5)  # left  hand 1
        pygame.draw.line(screen, black, [0 , 350], [1000 , 350], 5)  # road
        pygame.draw.line(screen, black, [0 , 450], [1000 , 450], 5)  # road

        k += 10
        return k

    done = False
    clock = pygame.time.Clock()

    while done == False:
        clock.tick(7)

        for event in pygame.event.get():  # Проходимся по событиям
            if event.type == pygame.QUIT:  # Если пользователь закрыл окно
                done = True  # Сигнализируем что цикл пора завершать

        screen.fill(white)

        if i == 1:
            k = line1(k)
        if i == 2:
            k = line2(k)
        if i == 3:
            k = line3(k)
            i = 0
        i+=1
        if k > 950:
            k = -100
        pygame.display.flip()

    pygame.quit()