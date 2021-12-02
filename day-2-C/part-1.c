#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Position {
  int x;
  int y;
};

struct InputLine { char* direction; int magnitude; };

struct InputLine split_string(char* line)
{
    struct InputLine curr = { .direction = NULL, .magnitude = 0 };
    char * pch = strtok (line," ,.-");
    curr.direction = pch;


    pch = strtok (NULL," ,.-");
    curr.magnitude = atoi(pch);

    return curr;
}


int main(void)
{
    FILE * fp;
    char * line = NULL;
    size_t len = 0;
    ssize_t read;

    fp = fopen("input.txt", "r");
    if (fp == NULL)
        exit(EXIT_FAILURE);

    struct Position position = { .x = 0, .y = 0 };

    int lines = 0;
    while ((read = getline(&line, &len, fp)) != -1) {
      lines++;
      struct InputLine curr = split_string(line);
      if (strcmp(curr.direction, "forward") == 0)
      {
        position.x += curr.magnitude;
      }
      else if (strcmp(curr.direction, "up") == 0)
      {
        position.y -= curr.magnitude;
      }
      else if (strcmp(curr.direction, "down") == 0)
      {
        position.y += curr.magnitude;
      }
    }

    fclose(fp);
    if (line)
        free(line);


    printf("There is %u lines in this file\n", lines);
    printf("X POSITON: %u, Y POSITION: %u\n", position.x, position.y);
    printf("%u", position.x * position.y);

    exit(EXIT_SUCCESS);
}
