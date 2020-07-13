"""
HackerNews module.
"""
import logging

import asyncio
from operator import itemgetter

import httpx

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

class HackerNews:
    """ Get the top posts trending on HackerNews. """

    def __init__(self):
        self.url = 'https://hacker-news.firebaseio.com/v0/topstories.json'
        self.container = {}
        self.stories = []
        self.responses = []
        self.run()

    def _resp(self):
        """ Return topstories.json which is just an array of post id's. """
        try:
            resp = httpx.get(self.url)
            return resp.json()
        except TimeoutError as err:
            logger.error(err)
        except asyncio.TimeoutError as err:
            logger.error(err)
        except httpx.ConnectTimeout as err:
            logger.error(err)

    def top_ten(self):
        """
        Return a list of the top ten HackerNews stories.
        :return: HackerNews list of story id's
        """
        top_ten = []
        for story_id in self._resp()[:10]:
            top_ten.append(story_id)
        return top_ten

    @staticmethod
    async def fetch(item_id, client):
        """
        Return async response object for HackerNews item.
        :param item_id: json id for HackerNews item
        :param client: httpx.AsyncClient method
        :return: httpx.AsyncClient response object
        """
        return await client.get(
            f"https://hacker-news.firebaseio.com/v0/item/{item_id}.json")

    async def story_metadata(self):
        """
        Return a list of responses from the given list of HackerNews item id's.
        """
        _list = self.top_ten()
        logger.debug(_list)
        async with httpx.AsyncClient() as client:
            logger.debug(client)
            self.responses = await asyncio.gather(
                *[self.fetch(id, client) for id in _list])

    def run_async_task(self, coro=None):
        """
        A helper class to create async coroutines.
        :param coro: a valid coroutine
        :return: result set from the async function
        """
        if coro is None:
            coro = self.story_metadata()

        logger.debug(coro)
        results = asyncio.run(coro)
        return results

    def extract_json(self):
        """
        Extract the response JSON and append to a list
        """
        for data in self.responses:
            results = self.container = data.json()
            self.stories.append(results)

    def sort_by_score(self):
        """
        Sort the list of HackerNews stories by score.
        """
        self.stories = sorted(self.stories, key=itemgetter('score'),
                              reverse=True)

    def run(self):
        """
        Entry point for the class, kicking off all the methods.
        """
        self.run_async_task(self.story_metadata())
        self.extract_json()
        self.sort_by_score()


hn = HackerNews()
hn.run()
# print(hn.responses)
from pprint import pprint
pprint(hn.stories)