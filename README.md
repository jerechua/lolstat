LoL Stat
========

To be able to view the stats matrix of a champion matchup, for both a
single player, and the entire player base.

Why would you want to use this?
-------------------------------

For example, you would use this to find the "counter" to the champion
you see picked on the other team, and you want to find the champion that
has the highest win rate against that other champion.

Roadblocks
----------

- I stopped playing the game, don't feel like continuing this project.
- Riot API is hard to work with
  - Caching all the stats is painful. Need to get "X" number of games per
    player, where "X" needs to be high enough otherwise it's pointless.
    This could be crawled, but will hit rate limit fairly quickly.
  - Rate limits are way too low especially for development.
  
What's in the code
------------------

- Weak attempts at creating a rate limiter and http libraries
- Riot API client (which could actually be useful for others)
- Caching version of the Riot API client, which queries DB first,
  and only requests if it doesn't exist. Uses Gorm in SQL for Users,
  and MGo in MongoDB for Match store.
