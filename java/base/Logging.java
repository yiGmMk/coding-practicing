package base;

import java.util.logging.Logger;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

/**
 * Logging
 */
public class Logging {

    static final Log log = LogFactory.getLog(Logging.class);

    public static void main(String[] args) {

        // 标准库中的log
        Logger logger = Logger.getGlobal();
        logger.info("start process...");
        logger.warning("memory is running out...");
        logger.fine("ignored.");
        logger.severe("process will be terminated...");

        // 外部包,手动导入需要下载相应的jar包并加入Referenced Libraries中
        log.info("start...");
        log.warn("end.");
    }
}