

import logging

logger = logging.getLogger(__name__)

class Notifier(object):
    def __init___(self):
        pass

    @classmethod
    def from_conf(cls, conf):
        """
        Create a notifier from the given conf

        @param conf: a conf object
        """
        return Notifier()


    def notify(self, window):
        """
        Notify that the window is anomalous
        """
        logger.info("Window: %s is anomalous" % window)
